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

func (s *Service) GetFriendRequests(username string) ([]*entity.User, error) {
	reqUsers, err := s.friendRepo.GetFriendRequests(username)
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (s *Service) CheckFriendRequest(senderName string, receiverName string) (bool, error) {
	return s.friendRepo.CheckFriendRequest(senderName, receiverName)
}

func (s *Service) SendFriendRequest(senderName string, receiverName string) (FriendRequestResult, error) {
	friendExists, err := s.friendRepo.CheckFriendship(senderName, receiverName)
	if err != nil {
		return FriendRequestNone, err
	}
	if friendExists {
		return FriendRequestNone, nil
	}

	friendReqExists, err := s.friendRepo.CheckFriendRequest(senderName, receiverName)
	if err != nil {
		return FriendRequestNone, err
	}
	if friendReqExists {
		return FriendRequestNone, nil
	}

	isReciprocal, err := s.friendRepo.CheckFriendRequest(receiverName, senderName)
	if err != nil {
		return FriendRequestNone, err
	}
	if isReciprocal {
		if err := s.friendRepo.DeleteFriendRequest(receiverName, senderName); err != nil {
			return FriendRequestNone, err
		}
		if err := s.friendRepo.AddFriend(senderName, receiverName); err != nil {
			return FriendRequestNone, err
		}
		return FriendRequestAccepted, nil
	} else {
		if err := s.friendRepo.SendFriendRequest(senderName, receiverName); err != nil {
			return FriendRequestNone, err
		}
		return FriendRequestSend, nil
	}
}

func (s *Service) DeleteFriendRequest(senderName string, receiverName string) error {
	if err := s.friendRepo.DeleteFriendRequest(senderName, receiverName); err != nil {
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

func (s *Service) CheckFriendship(userA string, userBName string) (bool, error) {
	return s.friendRepo.CheckFriendship(userA, userBName)
}

func (s *Service) DeleteFriend(userA string, userBName string) error {
	if err := s.friendRepo.DeleteFriend(userA, userBName); err != nil {
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
