package payload

type FriendshipType string

const (
	None     FriendshipType = "none"
	Friend   FriendshipType = "friend"
	Sent     FriendshipType = "sent"
	Received FriendshipType = "received"
)

type UserCreateRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserUpdateRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
}

type UserListRequest struct {
	Usernames []string `json:"usernames"`
}

type FriendshipResponse struct {
	Friendship FriendshipType `json:"friendship"`
}

type CheckFriendRequest struct {
	User   string `json:"username"`
	Friend string `json:"friend"`
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
