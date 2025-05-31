package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gateway/client"
	"gateway/internal"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[string]*Client)
var groupMembers = make(map[int][]string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MakeHandler(app *gin.Engine, groupClient client.GroupClient) {
	app.GET("/ws", func(c *gin.Context) {
		handleWebsocket(c, groupClient)
	})

	app.POST("/ws/message", func(c *gin.Context) {
		handleChatMessage(c, groupClient)
	})

	app.POST("/ws/noti", func(c *gin.Context) {
		handleNotificationMessage(c)
	})
}

func handleChatMessage(c *gin.Context, groupClient client.GroupClient) {
	username := c.Request.Header.Get("X-Username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No username provided"})
		return
	}

	var body ChatPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	body.Username = username

	fmt.Printf("[CHAT] %v: %v\n", len(clients), clients)
	if _, ok := clients[body.Username]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not connected"})
		return
	}

	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Failed to marshal chat payload:", err)
		return
	}
	msg := Message{
		Type:    MessageChat,
		Payload: json.RawMessage(payload),
	}

	if !tryBroadcastToGroup(body.Username, body.GroupID, msg) {
		_, err := updateGroups(groupClient, username, username)
		if err != nil {
			fmt.Println("Failed to update groups:", err)
			return
		}
		if !tryBroadcastToGroup(body.Username, body.GroupID, msg) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found or user not in group"})
			return
		}
	}
}

func handleNotificationMessage(c *gin.Context) {
	username := c.Request.Header.Get("X-Username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No username provided"})
		return
	}

	var body NotificationPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("[NOTI] %v: %v\n", len(clients), clients)
	client, ok := clients[body.Username]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not connected"})
		return
	}

	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Failed to marshal chat payload:", err)
		return
	}
	msg := Message{
		Type:    MessageNotification,
		Payload: json.RawMessage(payload),
	}
	if client.Conn.WriteJSON(msg) != nil {
		fmt.Println("Failed to send notification message through socket:", err)
		return
	}
}

func handleWebsocket(c *gin.Context, groupClient client.GroupClient) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}

	_, msgBytes, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("WebSocket read error during auth:", err)
		conn.Close()
		return
	}

	var authMsg Message
	if err := json.Unmarshal(msgBytes, &authMsg); err != nil {
		fmt.Println("Invalid message format during auth:", err)
		conn.Close()
		return
	}

	if authMsg.Type != MessageAuth {
		fmt.Println("First message must be auth type")
		conn.Close()
		return
	}

	var authPayload AuthPayload
	if err := json.Unmarshal(authMsg.Payload, &authPayload); err != nil {
		fmt.Println("Invalid auth payload:", err)
		conn.Close()
		return
	}

	username, err := internal.ParseJwt(authPayload.Token)
	if err != nil {
		fmt.Println("Invalid token:", err)
		conn.Close()
		return
	}

	groups, err := updateGroups(groupClient, username, username)
	if err != nil {
		fmt.Println("Failed to update groups:", err)
		conn.Close()
		return
	}

	gids := make([]int, len(groups))
	for i, g := range groups {
		gids[i] = g.ID
	}

	client := &Client{
		Username: username,
		Conn:     conn,
		Groups:   gids,
	}

	for _, group := range groups {
		if groupMembers[group.ID] == nil {
			groupMembers[group.ID] = make([]string, 1)
		}
		groupMembers[group.ID] = append(groupMembers[group.ID], username)
	}
	clients[username] = client

	fmt.Printf("[CONNECTED] %s in groups %v\n", username, groups)

	authSuccessMsg := Message{
		Type:    MessageAuth,
		Payload: json.RawMessage(`{"status":"success"}`),
	}
	authSuccessBytes, _ := json.Marshal(authSuccessMsg)
	conn.WriteMessage(websocket.TextMessage, authSuccessBytes)

	go handleMessages(client)
}

func handleMessages(client *Client) {
	defer func() {
		fmt.Printf("[DISCONNECTED] %s\n", client.Username)

		for _, groupID := range client.Groups {
			if members, exists := groupMembers[groupID]; exists {
				for i, member := range members {
					if member == client.Username {
						groupMembers[groupID] = append(members[:i], members[i+1:]...)
						break
					}
				}
				if len(members) == 0 {
					delete(groupMembers, groupID)
				}
			}
		}

		delete(clients, client.Username)
		client.Conn.Close()
	}()

	for {
		_, msgBytes, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			continue
		}

		switch msg.Type {
		case MessageAuth:
			fmt.Println("Ignoring additional auth message from authenticated client")
			continue

		case MessageChat:
			var chat ChatPayload
			if err := json.Unmarshal(msg.Payload, &chat); err != nil {
				fmt.Println("Invalid chat payload:", err)
				continue
			}
			broadcastToGroup(chat.GroupID, msgBytes, chat.Username)

		case MessageNotification:
			var notif NotificationPayload
			if err := json.Unmarshal(msg.Payload, &notif); err != nil {
				fmt.Println("Invalid notification payload:", err)
				continue
			}

			fmt.Printf("[NOTIFICATION] To %s: %s - %s\n", client.Username, notif.Desc, notif.Link)
			client.Conn.WriteMessage(websocket.TextMessage, msgBytes)

		default:
			fmt.Println("Unknown message type:\n", msg)
		}
	}
}

func updateGroups(groupClient client.GroupClient, authUsername string, username string) ([]*client.Group, error) {
	gu, err := groupClient.GetGroupsOfUser(authUsername, username)
	if err != nil {
		return nil, err
	}
	return gu, nil
}

func broadcastToGroup(groupID int, message []byte, ignore string) {
	if users, ok := groupMembers[groupID]; ok {
		for _, username := range users {
			if username == ignore {
				continue
			}
			if client, ok := clients[username]; ok {
				client.Conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}

func tryBroadcastToGroup(ignore string, groupID int, message interface{}) bool {
	users, ok := groupMembers[groupID]
	if ok {
		for _, username := range users {
			if username == ignore {
				continue
			}
			if client, ok := clients[username]; ok {
				client.Conn.WriteJSON(message)
			}
		}
	}
	return ok
}
