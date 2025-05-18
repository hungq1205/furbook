package group

import (
	"message/entity"
	"message/util"
)

type UseCase interface {
	GetGroup(groupID int) (*entity.Group, error)
	UpdateGroup(groupID int, groupName string) (*entity.Group, error)
	CreateGroup(ownername string, groupName string, members []string) (*entity.Group, error)
	DeleteGroup(groupID int) error

	GetGroupsOfUser(username string, pagination util.Pagination) ([]*entity.Group, error)
	CheckOwnership(username string, groupID int) (bool, error)
	CheckMembership(username string, groupID int) (bool, error)
	GetMembers(groupID int) ([]string, error)
	AddMember(groupID int, username string) (*entity.Group, error)
	RemoveMember(groupID int, username string) (*entity.Group, error)
}
