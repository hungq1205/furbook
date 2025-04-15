package friend

import (
	"user-service/entity"
	"user-service/infrastructure/repository/friend"
)

type FriendService struct {
	friendRepo *friend.FriendRepository
}

func (s *FriendService) GetFriendRequests(username string) ([]*entity.User, error) {
	reqUsers, err := s.friendRepo.GetFriendRequests(username)
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (s *FriendService) CheckFriendRequest(sender string, receiver string) (bool, error) {
	return s.friendRepo.CheckFriendRequest(sender, receiver)
}

func (s *FriendService) SendFriendRequest(sender string, receiver string) error {
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

func (s *FriendService) DeleteFriendRequest(sender string, receiver string) error {
	if err := s.friendRepo.DeleteFriendRequest(sender, receiver); err != nil {
		return err
	}
	return nil
}

func (s *FriendService) CountFriendRequests(username string) (int, error) {
	count, err := s.friendRepo.CountFriendRequests(username)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *FriendService) GetFriends(username string) ([]*entity.User, error) {
	friends, err := s.friendRepo.GetFriends(username)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (s *FriendService) CheckFriendship(usernA string, userB string) (bool, error) {
	return s.friendRepo.CheckFriendship(usernA, userB)
}

func (s *FriendService) DeleteFriend(sender string, receiver string) error {
	if err := s.friendRepo.DeleteFriend(sender, receiver); err != nil {
		return err
	}
	return nil
}

func (s *FriendService) CountFriends(username string) (int, error) {
	count, err := s.friendRepo.CountFriends(username)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
