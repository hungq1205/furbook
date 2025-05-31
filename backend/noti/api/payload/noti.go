package payload

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

type NotiUpdateRequest struct {
	Read bool `json:"read"`
}
