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
	db.AutoMigrate(&entity.Message{})
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message *entity.Message) (*entity.Message, error) {
	err := r.db.Model(&entity.Message{}).Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessageRepository) DeleteMessagesByUser(username string) error {
	err := r.db.Model(&entity.Message{}).Where("username = ?", username).Delete(&entity.Message{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetGroupMessageList(groupID int, pagination util.Pagination) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Model(&entity.Message{}).
		Joins("join groups g on g.id = messages.group_id").
		Where("group_id = ?", groupID).
		Find(&messages).
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetDirectMessageList(userA string, userB string, pagination util.Pagination) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.
		Model(&entity.Message{}).
		Joins("join group_users ug1 on ug1.username = ?", userA).
		Joins("join group_users ug2 on ug2.username = ? and ug2.group_id = ug1.group_id", userB).
		Joins("join groups on groups.id = messages.group_id").
		Where("groups.is_direct = ?", true).
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Find(&messages).
		Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetLastMessage(groupID int) (*entity.Message, error) {
	var message entity.Message
	err := r.db.
		Model(&entity.Message{}).
		Where("group_id = ?", groupID).
		// Order("created_at desc").
		Last(&message).
		Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}
