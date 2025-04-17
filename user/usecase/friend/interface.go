package friend

import "user/entity"

type UseCase interface {
	GetFriendRequests(username string) ([]*entity.User, error)
	CheckFriendRequest(sender string, receiver string) (bool, error)
	SendFriendRequest(sender string, receiver string) error
	DeleteFriendRequest(sender string, receiver string) error
	CountFriendRequests(username string) (int, error)

	GetFriends(username string) ([]*entity.User, error)
	DeleteFriend(sender string, receiver string) error
	CheckFriendship(usernA string, userB string) (bool, error)
	CountFriends(username string) (int, error)
}
