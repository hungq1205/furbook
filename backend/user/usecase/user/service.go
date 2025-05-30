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

func (s *Service) GetUser(username string) (*entity.User, error) {
	usr, err := s.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) GetUsers(usernames []string) ([]*entity.User, error) {
	users, err := s.userRepo.GetUsers(usernames)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) CreateUser(username string, displayName string) (*entity.User, error) {
	usr, err := s.userRepo.CreateUser(username, displayName)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) UpdateUser(username string, avatar string, bio string) (*entity.User, error) {
	usr, err := s.userRepo.UpdateUser(username, avatar, bio)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) DeleteUser(username string) error {
	if err := s.userRepo.DeleteUser(username); err != nil {
		return err
	}
	return nil
}
