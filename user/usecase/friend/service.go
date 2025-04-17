package friend

import (
	"user/entity"
	"user/infrastructure/repository/friend"
)

type Service struct {
	friendRepo *friend.FriendRepository
}

func NewService(friendRepo *friend.FriendRepository) *Service {
	return &Service{
		friendRepo: friendRepo,
	}
}

func (s *Service) GetFriendRequests(username string) ([]*entity.User, error) {
	reqUsers, err := s.friendRepo.GetFriendRequests(username)
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (s *Service) CheckFriendRequest(sender string, receiver string) (bool, error) {
	return s.friendRepo.CheckFriendRequest(sender, receiver)
}

func (s *Service) SendFriendRequest(sender string, receiver string) error {
	friendExists, err := s.friendRepo.CheckFriendship(sender, receiver)
	if err != nil {
		return err
	}
	if friendExists {
		return nil
	}

	friendReqExists, err := s.friendRepo.CheckFriendRequest(sender, receiver)
	if err != nil {
		return err
	}
	if friendReqExists {
		return nil
	}

	isReciprocal, err := s.friendRepo.CheckFriendRequest(receiver, sender)
	if err != nil {
		return err
	}
	if isReciprocal {
		if err := s.friendRepo.DeleteFriendRequest(receiver, sender); err != nil {
			return err
		}
		if err := s.friendRepo.AddFriend(sender, receiver); err != nil {
			return err
		}
	} else {
		if err := s.friendRepo.SendFriendRequest(sender, receiver); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) DeleteFriendRequest(sender string, receiver string) error {
	if err := s.friendRepo.DeleteFriendRequest(sender, receiver); err != nil {
		return err
	}
	return nil
}

func (s *Service) CountFriendRequests(username string) (int, error) {
	count, err := s.friendRepo.CountFriendRequests(username)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *Service) GetFriends(username string) ([]*entity.User, error) {
	friends, err := s.friendRepo.GetFriends(username)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (s *Service) CheckFriendship(usernA string, userB string) (bool, error) {
	return s.friendRepo.CheckFriendship(usernA, userB)
}

func (s *Service) DeleteFriend(sender string, receiver string) error {
	if err := s.friendRepo.DeleteFriend(sender, receiver); err != nil {
		return err
	}
	return nil
}

func (s *Service) CountFriends(username string) (int, error) {
	count, err := s.friendRepo.CountFriends(username)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
