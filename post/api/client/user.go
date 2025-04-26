package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"post/api/presenter"
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

type GetUserListResponse struct {
	Users []*presenter.User `json:"users"`
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

	resp, err := http.Post(c.userUrl+"/api/users", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody GetUserListResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody.Users, nil
}
