package noti

import (
	"noti/entity"
	"noti/util"
)

type UseCase interface {
	GetNoti(id int) (*entity.Notification, error)
	GetNotisOfUser(username string, pagination util.Pagination) ([]*entity.Notification, error)
	CreateNoti(username, icon, desc, link string) (*entity.Notification, error)
	CreateNotiToUsers(usernames []string, icon, desc, link string) ([]*entity.Notification, error)
	UpdateNoti(id int, read bool) (*entity.Notification, error)
	DeleteNoti(id int) error
}
