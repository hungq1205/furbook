package user

import "user/entity"

type UseCase interface {
	GetUser(userID uint) (*entity.User, error)
	GetUsers(userIDs []uint) ([]*entity.User, error)
	CreateUser(username string, avatar string) (*entity.User, error)
	UpdateUser(userID uint, avatar string) (*entity.User, error)
	DeleteUser(userID uint) error
}
