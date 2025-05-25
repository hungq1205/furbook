package client

import (
	"encoding/json"
	"net/http"
)

type GroupClient interface {
	FindDirectGroup(username, oppUsername string) (int, error)
}

type GroupClientImpl struct {
	groupUrl string
}

func NewGroupClient(groupUrl string) GroupClient {
	return &GroupClientImpl{
		groupUrl: groupUrl,
	}
}

func (c *GroupClientImpl) FindDirectGroup(username, oppUsername string) (int, error) {
	req, err := http.NewRequest("GET", c.groupUrl+"/api/group/direct/"+oppUsername, nil)
	if err != nil {
		return -1, err
	}
	req.Header.Set("X-Username", username)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	var group struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return -1, err
	}

	return group.ID, nil
}
