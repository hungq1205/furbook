package group

import (
	"message/entity"
	"message/infrastructure/repository"
	"message/util"
)

type Service struct {
	groupRepo     *repository.GroupRepository
	groupUserRepo *repository.GroupUserRepository
}

func NewService(groupRepo *repository.GroupRepository, groupUserRepo *repository.GroupUserRepository) *Service {
	return &Service{
		groupRepo:     groupRepo,
		groupUserRepo: groupUserRepo,
	}
}

func (s *Service) GetGroup(groupID int) (*entity.Group, error) {
	group, err := s.groupRepo.GetGroup(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (s *Service) UpdateGroup(groupID int, groupName string) (*entity.Group, error) {
	group, err := s.groupRepo.UpdateGroup(&entity.Group{ID: groupID, Name: groupName})
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (s *Service) CreateGroup(ownername string, groupName string, members []string) (*entity.Group, error) {
	group, err := s.groupRepo.CreateGroup(&entity.Group{OwnerName: ownername, Name: groupName})
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		err = s.groupUserRepo.AddUserToGroup(group.ID, member)
		if err != nil {
			return nil, err
		}
	}
	return group, nil
}

func (s *Service) DeleteGroup(groupID int) error {
	err := s.groupRepo.DeleteGroup(groupID)
	if err != nil {
		return err
	}

	err = s.groupUserRepo.RemoveUsersFromGroup(groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetGroupsOfUser(username string, pagination util.Pagination) ([]*entity.Group, error) {
	users, err := s.groupUserRepo.GetGroupsOfUser(username, pagination)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) CheckOwnership(username string, groupID int) (bool, error) {
	group, err := s.groupRepo.GetGroup(groupID)
	if err != nil {
		return false, err
	}
	return group.OwnerName == username, nil
}

func (s *Service) CheckMembership(username string, groupID int) (bool, error) {
	return s.groupUserRepo.CheckUserInGroup(groupID, username)
}

func (s *Service) GetMembers(groupID int) ([]string, error) {
	users, err := s.groupUserRepo.GetUsersInGroup(groupID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) AddMember(groupID int, username string) (*entity.Group, error) {
	err := s.groupUserRepo.AddUserToGroup(groupID, username)
	if err != nil {
		return nil, err
	}
	return s.GetGroup(groupID)
}

func (s *Service) RemoveMember(groupID int, username string) (*entity.Group, error) {
	err := s.groupUserRepo.RemoveUserFromGroup(groupID, username)
	if err != nil {
		return nil, err
	}
	return s.GetGroup(groupID)
}
