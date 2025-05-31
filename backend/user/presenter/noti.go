package presenter

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Icon      string    `json:"icon"`
	Desc      string    `json:"desc"`
	Link      string    `json:"link"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

type NotiCreateRequest struct {
	Username string `json:"username"`
	Icon     string `json:"icon"`
	Desc     string `json:"desc"`
	Link     string `json:"link"`
}

type NotiToUsersCreateRequest struct {
	Usernames []string `json:"usernames"`
	Icon      string   `json:"icon"`
	Desc      string   `json:"desc"`
	Link      string   `json:"link"`
}
