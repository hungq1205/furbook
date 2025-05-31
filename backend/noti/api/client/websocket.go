package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"noti/entity"
)

type WsClient interface {
	SendNoti(noti *entity.Notification) error
}

type WsClientImpl struct {
	wsUrl string
}

func NewWsClient(wsUrl string) WsClient {
	return &WsClientImpl{
		wsUrl: wsUrl,
	}
}

func (c *WsClientImpl) SendNoti(noti *entity.Notification) error {
	body, err := json.Marshal(noti)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.wsUrl+"/ws/noti", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Username", "system")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
