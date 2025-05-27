package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Group struct {
	ID        int    `json:"id"`
	IsDirect  bool   `json:"is_direct"`
	OwnerName string `json:"owner_name"`
}

type GroupClient interface {
	GetGroup(string, int) (*Group, error)
	GetGroupsOfUser(string, string) ([]*Group, error)
}

type GroupClientImpl struct {
	groupUrl string
}

func NewGroupClient(groupUrl string) GroupClient {
	return &GroupClientImpl{
		groupUrl: groupUrl,
	}
}

func (c *GroupClientImpl) GetGroup(authUsername string, id int) (*Group, error) {
	req, err := http.NewRequest("GET", c.groupUrl+"/api/group/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Username", authUsername)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get group: %s", resp.Status)
	}

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, err
	}

	return &group, nil
}

func (c *GroupClientImpl) GetGroupsOfUser(authUsername string, username string) ([]*Group, error) {
	req, err := http.NewRequest("GET", c.groupUrl+"/api/group?username="+username, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Username", authUsername)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get groups of user: %s", resp.Status)
	}

	var groups []*Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, err
	}

	return groups, nil
}
