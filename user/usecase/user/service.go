package user

import (
	"user/entity"
	"user/infrastructure/repository/user"
)

type Service struct {
	userRepo *user.UserRepository
}

func NewService(userRepo *user.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) GetUser(userID uint) (*entity.User, error) {
	usr, err := s.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) GetUsers(userIDs []uint) ([]*entity.User, error) {
	users, err := s.userRepo.GetUsers(userIDs)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) CreateUser(username string, avatar string) (*entity.User, error) {
	usr, err := s.userRepo.CreateUser(username, avatar)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) UpdateUser(userID uint, avatar string, bio string) (*entity.User, error) {
	usr, err := s.userRepo.UpdateUser(userID, avatar, bio)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) DeleteUser(userID uint) error {
	if err := s.userRepo.DeleteUser(userID); err != nil {
		return err
	}
	return nil
}
