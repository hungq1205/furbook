package presenter

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	FriendNum   int    `json:"friendNum"`
	GroupId     *int   `json:"groupid"`
}

type FriendList struct {
	Friends []*User `json:"friends"`
}
