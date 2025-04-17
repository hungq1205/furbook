package user

import "user/entity"

type UseCase interface {
	GetUser(username string) (*entity.User, error)
	GetUsers(usernames []string) ([]*entity.User, error)
	CheckUsernameExists(username string) (bool, error)
	CreateUser(username string, avatar string) (*entity.User, error)
	UpdateUser(username string, avatar string) (*entity.User, error)
	DeleteUser(username string) error
}
