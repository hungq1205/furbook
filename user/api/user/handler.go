package user

import (
	"user-service/usecase/friend"
	"user-service/usecase/user"
)

type UserHandler struct {
	userService   user.UseCase
	friendService friend.UseCase
}

func NewUserHandler(userService user.UseCase, friendService friend.UseCase) *UserHandler {
	return &UserHandler{
		userService:   userService,
		friendService: friendService,
	}
}
