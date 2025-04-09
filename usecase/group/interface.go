package usecase

import (
	"context"
	"test/entity"
)

type UseCase interface {
	GetGroup(ctx context.Context, groupID int) (*entity.Group, error)
	UpdateGroup(ctx context.Context, groupID int, name string) (*entity.Group, error)
	CreateGroup(ctx context.Context, name string) (*entity.Group, error)
	DeleteGroup(ctx context.Context, groupID int) error

	GetGroupsOfUser(ctx context.Context) ([]*entity.Group, error)
	CheckOwnership(ctx context.Context, groupID int) (bool, error)
	GetMembers(ctx context.Context, groupID int) ([]*entity.Group, error)
	AddMember(ctx context.Context, groupID int, username string) (*entity.Group, error)
	RemoveMember(ctx context.Context, groupID int, username string) error
}
