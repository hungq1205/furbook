package websocket

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type MessageType string

const (
	MessageAuth         MessageType = "auth"
	MessageChat         MessageType = "chat"
	MessageNotification MessageType = "notification"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
	Groups   []int
}

type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type AuthPayload struct {
	Token string `json:"token"`
}

type ChatPayload struct {
	MessageID int       `json:"messageId"`
	Username  string    `json:"username"`
	GroupID   int       `json:"groupId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type NotificationPayload struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Icon      string    `json:"icon"`
	Desc      string    `json:"desc"`
	Link      string    `json:"link"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}
