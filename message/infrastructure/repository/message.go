package repository

import (
	"message/entity"
	"message/util"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message *entity.Message) (*entity.Message, error) {
	err := r.db.Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessageRepository) DeleteMessagesByUser(username string) error {
	err := r.db.Where("username = ?", username).Delete(&entity.Message{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetGroupMessageList(groupID int, pagination util.Pagination) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Joins("join groups g on g.id = messages.group_id").
		Where("group_id = ?", groupID).
		Find(&messages).
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *MessageRepository) GetDirectMessageList(userA string, userB string, pagination util.Pagination) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Joins("join user_groups ug1 on ug1.username = ?", userA).
		Joins("join user_groups ug2 on ug2.username = ?", userB).
		Where("groups.is_direct = ?", true).
		Find(&messages).
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
