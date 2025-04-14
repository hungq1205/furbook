package message

import (
	"errors"
	"gorm.io/gorm"
	"test/entity"
	"test/infrastructure/repository"
	"test/util"
)

type Service struct {
	messageRepo   *repository.MessageRepository
	groupUserRepo *repository.GroupUserRepository
}

func NewService(messageRepo *repository.MessageRepository, groupUserRepo *repository.GroupUserRepository) *Service {
	return &Service{
		messageRepo:   messageRepo,
		groupUserRepo: groupUserRepo,
	}
}

func (s *Service) SendMessage(username string, content string, groupID int) (*entity.Message, error) {
	msg, err := s.messageRepo.CreateMessage(&entity.Message{
		Username: username,
		Content:  content,
		GroupID:  groupID,
	})
	if err != nil {
		return nil, err
	}

	return msg, err
}

func (s *Service) SendDirectMessage(username string, oppUsername string, content string) (*entity.Message, error) {
	group, err := s.groupUserRepo.GetDirectGroup(username, oppUsername)
	if err != nil {
		return nil, err
	}
	groupID := group.ID

	msg, err := s.messageRepo.CreateMessage(&entity.Message{
		Username: username,
		Content:  content,
		GroupID:  groupID,
	})
	if err != nil {
		return nil, err
	}

	return msg, err
}

func (s *Service) GetDirectMessageList(username string, oppUsername string, pagination *util.Pagination) ([]*entity.Message, error) {
	messages, err := s.messageRepo.GetDirectMessageList(username, oppUsername, pagination)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*entity.Message{}, nil
	}
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *Service) GetGroupMessageList(groupID int, pagination *util.Pagination) ([]*entity.Message, error) {
	messages, err := s.messageRepo.GetGroupMessageList(groupID, pagination)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
