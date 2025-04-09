package usecase

import (
	"context"
	"test/entity"
)

type UseCase interface {
	SendMessage(ctx context.Context, content string, groupID int) (*entity.Message, error)
	SendDirectMessage(ctx context.Context, oppUsername string, content string) (*entity.Message, error)
	GetDirectMessageList(ctx context.Context, oppUsername string) ([]*entity.Message, error)
	GetGroupMessageList(ctx context.Context, groupID int) ([]*entity.Message, error)
}
