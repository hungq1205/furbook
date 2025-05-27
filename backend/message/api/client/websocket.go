package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type WsClient interface {
	SendMessage(int, string, int, string, time.Time) error
}

type WsClientImpl struct {
	wsUrl string
}

type WsMessageRequest struct {
	MessageID int       `json:"messageId"`
	GroupID   int       `json:"groupId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewWsClient(wsUrl string) WsClient {
	return &WsClientImpl{
		wsUrl: wsUrl,
	}
}

func (c *WsClientImpl) SendMessage(messageId int, username string, groupId int, content string, createdAt time.Time) error {
	body, err := json.Marshal(&WsMessageRequest{
		MessageID: messageId,
		GroupID:   groupId,
		Content:   content,
		CreatedAt: createdAt,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.wsUrl+"/ws/message", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Username", username)

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
