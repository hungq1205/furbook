package repository

import (
	"gorm.io/gorm"
	"test/entity"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) GetGroupMessageList(groupID int) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Joins("join groups g on g.id = messages.group_id").
		Where("group_id = ?", groupID).
		Find(&messages).
		Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *MessageRepository) GetDirectMessageList(userA string, userB string) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Joins("join user_groups ug1 on ug1.username = ?", userA).
		Joins("join user_groups ug2 on ug2.username = ?", userB).
		Where("groups.is_direct = ?", true).
		Find(&messages).
		Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
