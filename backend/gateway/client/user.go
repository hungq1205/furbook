package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	FriendNum   int    `json:"friendNum"`
}

type UserCreateRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type UserClient interface {
	GetUser(string) (*User, error)
	CreateUser(string, string) (*User, error)
}

type UserClientImpl struct {
	userUrl string
}

func NewUserClient(userUrl string) UserClient {
	return &UserClientImpl{
		userUrl: userUrl,
	}
}

func (c *UserClientImpl) GetUser(username string) (*User, error) {
	resp, err := http.Get(c.userUrl + "/api/user/" + username)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserClientImpl) CreateUser(username string, displayName string) (*User, error) {
	body, err := json.Marshal(&UserCreateRequest{Username: username, DisplayName: displayName})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.userUrl+"/api/user", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK || resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
