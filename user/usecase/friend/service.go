package friend

import (
	"user/entity"
	"user/infrastructure/repository/friend"
)

type Service struct {
	friendRepo *friend.Repository
}

func NewService(friendRepo *friend.Repository) *Service {
	return &Service{
		friendRepo: friendRepo,
	}
}

func (s *Service) GetFriendRequests(userID uint) ([]*entity.User, error) {
	reqUsers, err := s.friendRepo.GetFriendRequests(userID)
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (s *Service) CheckFriendRequest(senderID uint, receiverID uint) (bool, error) {
	return s.friendRepo.CheckFriendRequest(senderID, receiverID)
}

func (s *Service) SendFriendRequest(senderID uint, receiverID uint) error {
	friendExists, err := s.friendRepo.CheckFriendship(senderID, receiverID)
	if err != nil {
		return err
	}
	if friendExists {
		return nil
	}

	friendReqExists, err := s.friendRepo.CheckFriendRequest(senderID, receiverID)
	if err != nil {
		return err
	}
	if friendReqExists {
		return nil
	}

	isReciprocal, err := s.friendRepo.CheckFriendRequest(receiverID, senderID)
	if err != nil {
		return err
	}
	if isReciprocal {
		if err := s.friendRepo.DeleteFriendRequest(receiverID, senderID); err != nil {
			return err
		}
		if err := s.friendRepo.AddFriend(senderID, receiverID); err != nil {
			return err
		}
	} else {
		if err := s.friendRepo.SendFriendRequest(senderID, receiverID); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) DeleteFriendRequest(senderID uint, receiverID uint) error {
	if err := s.friendRepo.DeleteFriendRequest(senderID, receiverID); err != nil {
		return err
	}
	return nil
}

func (s *Service) CountFriendRequests(userID uint) (int, error) {
	count, err := s.friendRepo.CountFriendRequests(userID)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *Service) GetFriends(userID uint) ([]*entity.User, error) {
	friends, err := s.friendRepo.GetFriends(userID)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (s *Service) CheckFriendship(userAID uint, userBID uint) (bool, error) {
	return s.friendRepo.CheckFriendship(userAID, userBID)
}

func (s *Service) DeleteFriend(userAID uint, userBID uint) error {
	if err := s.friendRepo.DeleteFriend(userAID, userBID); err != nil {
		return err
	}
	return nil
}

func (s *Service) CountFriends(userID uint) (int, error) {
	count, err := s.friendRepo.CountFriends(userID)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
