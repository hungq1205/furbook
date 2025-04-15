package payload

type UserCreateRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserUpdateRequest struct {
	Avatar string `json:"avatar"`
}

type UserListRequest struct {
	Usernames []string `json:"usernames"`
}

type CheckFriendRequest struct {
	Username string `json:"username"`
	Friend   string `json:"friend"`
}

type ReceiverWrapper struct {
	Receiver string `json:"receiver"`
}

type FriendWrapper struct {
	Friend string `json:"friend"`
}

type SenderWrapper struct {
	Sender string `json:"sender"`
}
