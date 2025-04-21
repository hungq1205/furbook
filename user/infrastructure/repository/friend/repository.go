package friend

import (
	"errors"
	"user-service/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) GetFriends(userID uint) ([]*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Friends").Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*entity.User{}, nil
		}
		return nil, err
	}
	return user.Friends, nil
}

func (r *Repository) GetFriendUserIDs(userID uint) ([]uint, error) {
	friendIDs := []uint{}
	if err := r.db.
		Table("friendship").
		Where("user_id = ?", userID).
		Pluck("friend_id", &friendIDs).
		Error; err != nil {
		return nil, err
	}
	return friendIDs, nil
}

func (r *Repository) GetFriendRequests(userID uint) ([]*entity.User, error) {
	var reqUsers []*entity.User
	err := r.db.
		Table("friend_request fr").
		Select("u.*").
		Joins("JOIN user u ON fr.sender_id = u.id").
		Where("fr.receiver_id = ?", userID).
		Find(&reqUsers).Error
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (r *Repository) CountFriendRequests(userID uint) (int64, error) {
	var count int64
	if err := r.db.Table("friend_request").
		Where("receiver_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) CountFriends(userID uint) (int64, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) CheckFriendship(userAID uint, userBID uint) (bool, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("user_id = ? AND friend_id = ?", userAID, userBID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) CheckFriendRequest(senderID uint, receiverID uint) (bool, error) {
	var count int64
	if err := r.db.Table("friend_request").
		Where("sender_id = ? AND receiver_id = ?", senderID, receiverID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) AddFriend(userID uint, friendID uint) error {
	if err := r.db.Exec("INSERT INTO friendship (user_id, friend_id) VALUES (?, ?), (?, ?)",
		userID, friendID, friendID, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFriend(userID uint, friendID uint) error {
	if err := r.db.Exec("DELETE FROM friendship WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		userID, friendID, friendID, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) SendFriendRequest(senderID uint, receiverID uint) error {
	if err := r.db.Exec("INSERT INTO friend_request (sender_id, receiver_id) VALUES (?, ?)", senderID, receiverID).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFriendRequest(senderID uint, receiverID uint) error {
	if err := r.db.Exec("DELETE FROM friend_request WHERE sender_id = ? AND receiver_id = ?", senderID, receiverID).Error; err != nil {
		return err
	}
	return nil
}
