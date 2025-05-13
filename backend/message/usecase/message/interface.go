package message

import (
	"message/entity"
	"message/util"
)

type UseCase interface {
	SendMessage(username string, content string, groupID int) (*entity.Message, error)
	SendDirectMessage(username string, oppUsername string, content string) (*entity.Message, error)
	GetDirectMessageList(username string, oppUsername string, pagination util.Pagination) ([]*entity.Message, error)
	GetGroupMessageList(groupID int, pagination util.Pagination) ([]*entity.Message, error)
}
