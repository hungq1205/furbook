package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"user/presenter"
)

type NotiClient interface {
	CreateNoti(username, icon, desc, link string) (*presenter.Notification, error)
	CreateNotiToUsers(usernames []string, icon, desc, link string) error
}

type NotiClientImpl struct {
	notiUrl string
}

func NewNotiClient(notiUrl string) NotiClient {
	return &NotiClientImpl{
		notiUrl: notiUrl,
	}
}

func (c *NotiClientImpl) CreateNoti(username, icon, desc, link string) (*presenter.Notification, error) {
	body, err := json.Marshal(&presenter.NotiCreateRequest{
		Username: username,
		Icon:     icon,
		Desc:     desc,
		Link:     link,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.notiUrl+"/api/noti", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Username", "system")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var noti presenter.Notification
	if err := json.NewDecoder(resp.Body).Decode(&noti); err != nil {
		return nil, err
	}

	return &noti, nil
}

func (c *NotiClientImpl) CreateNotiToUsers(usernames []string, icon, desc, link string) error {
	body, err := json.Marshal(&presenter.NotiToUsersCreateRequest{
		Usernames: usernames,
		Icon:      icon,
		Desc:      desc,
		Link:      link,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.notiUrl+"/api/noti/createMultiple", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Username", "system")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
