package friend

import "user/entity"

type FriendRequestResult int

const (
	FriendRequestNone FriendRequestResult = iota
	FriendRequestSend
	FriendRequestAccepted
)

type UseCase interface {
	GetFriendRequests(username string) ([]*entity.User, error)
	CheckFriendRequest(senderName string, receiverName string) (bool, error)
	SendFriendRequest(senderName string, receiverName string) (FriendRequestResult, error)
	DeleteFriendRequest(senderName string, receiverName string) error
	CountFriendRequests(username string) (int, error)

	GetFriends(username string) ([]*entity.User, error)
	DeleteFriend(userA string, userBName string) error
	CheckFriendship(userA string, userBName string) (bool, error)
	CountFriends(username string) (int, error)
}
