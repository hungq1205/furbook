package friend

import "user-service/entity"

type UseCase interface {
	GetFriendRequests(userID uint) ([]*entity.FriendRequest, error)
	CheckFriendRequest(senderID uint, receiverID uint) (bool, error)
	SendFriendRequest(senderID uint, receiverID uint) error
	DeleteFriendRequest(senderID uint, receiverID uint) error
	CountFriendRequests(userID uint) (int64, error)

	GetFriends(userID uint) ([]*entity.User, error)
	DeleteFriend(userAID uint, userBID uint) error
	CheckFriendship(userAID uint, userBID uint) (bool, error)
	CountFriends(userID uint) (int, error)
}
