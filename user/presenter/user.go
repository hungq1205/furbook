package presenter

type User struct {
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	FriendNum int    `json:"friend_num"`
}

type FriendList struct {
	Friends []*User `json:"friends"`
}
