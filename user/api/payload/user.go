package payload

type UserCreateRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserUpdateRequest struct {
	UserID uint   `json:"user_id"`
	Avatar string `json:"avatar"`
}

type UserListRequest struct {
	UserIDs []uint `json:"user_ids"`
}

type CheckFriendRequest struct {
	UserID   uint `json:"user_id"`
	FriendID uint `json:"friend_id"`
}

type ReceiverWrapper struct {
	ReceiverID uint `json:"receiver_id"`
}

type FriendWrapper struct {
	FriendID uint `json:"friend_id"`
}

type SenderWrapper struct {
	SenderID uint `json:"sender_id"`
}
