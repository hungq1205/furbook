package user

import "user/entity"

type UseCase interface {
	GetUser(username string) (*entity.User, error)
	GetUsers(usernames []string) ([]*entity.User, error)
	CreateUser(username string, displayName string) (*entity.User, error)
	UpdateUser(username string, avatar string, bio string) (*entity.User, error)
	DeleteUser(username string) error
}
