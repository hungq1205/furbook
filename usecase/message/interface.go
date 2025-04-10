package message

import (
	"test/entity"
)

type UseCase interface {
	SendMessage(username string, content string, groupID int) (*entity.Message, error)
	SendDirectMessage(username string, oppUsername string, content string) (*entity.Message, error)
	GetDirectMessageList(username string, oppUsername string) ([]*entity.Message, error)
	GetGroupMessageList(groupID int) ([]*entity.Message, error)
}
