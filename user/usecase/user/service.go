package user

import (
	"user-service/entity"
	"user-service/infrastructure/repository/user"
)

type UserService struct {
	userRepo *user.UserRepository
}

func (s *UserService) GetUser(username string) (*entity.User, error) {
	user, err := s.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsers(usernames []string) ([]*entity.User, error) {
	users, err := s.userRepo.GetUsers(usernames)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CheckUsernameExists(username string) (bool, error) {
	exists, err := s.userRepo.CheckUsernameExists(username)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) CreateUser(username string, avatar string) (*entity.User, error) {
	user, err := s.userRepo.CreateUser(username, avatar)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(username string, avatar string) (*entity.User, error) {
	user, err := s.userRepo.UpdateUser(username, avatar)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(username string) error {
	if err := s.userRepo.DeleteUser(username); err != nil {
		return err
	}
	return nil
}
