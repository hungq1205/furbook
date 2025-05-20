package client

import (
	"bytes"
	"encoding/json"
	"message/api/presenter"
	"net/http"
)

type UserClient interface {
	FindUsers([]string) ([]*presenter.User, error)
}

type UserClientImpl struct {
	userUrl string
}

type GetUserListRequest struct {
	Usernames []string `json:"usernames"`
}

func NewUserClient(userUrl string) UserClient {
	return &UserClientImpl{
		userUrl: userUrl,
	}
}

func (c *UserClientImpl) FindUsers(usernames []string) ([]*presenter.User, error) {
	body, err := json.Marshal(&GetUserListRequest{Usernames: usernames})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.userUrl+"/api/user", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []*presenter.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
