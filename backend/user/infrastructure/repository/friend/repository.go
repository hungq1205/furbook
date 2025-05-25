package friend

import (
	"errors"
	"user/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&entity.User{}, &entity.FriendRequest{})
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetFriends(username string) ([]*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Friends").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*entity.User{}, nil
		}
		return nil, err
	}
	return user.Friends, nil
}

func (r *Repository) GetFriendUsernames(username string) ([]string, error) {
	friendUsernames := []string{}
	if err := r.db.
		Table("friendship f").
		Select("f.friend_name").
		Where("f.username = ?", username).
		Pluck("friend_name", &friendUsernames).
		Error; err != nil {
		return nil, err
	}
	return friendUsernames, nil
}

func (r *Repository) GetFriendRequests(username string) ([]*entity.User, error) {
	var reqUsers []*entity.User
	err := r.db.
		Table("friend_requests fr").
		Select("u.*").
		Joins("JOIN user u ON fr.sender = u.username").
		Where("fr.receiver = ?", username).
		Find(&reqUsers).Error
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (r *Repository) CountFriendRequests(username string) (int64, error) {
	var count int64
	if err := r.db.Table("friend_requests").
		Where("receiver = ?", username).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) CountFriends(username string) (int64, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) CheckFriendship(userA string, userB string) (bool, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("username = ? AND friend_name = ?", userA, userB).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) CheckFriendRequest(senderName string, receiverName string) (bool, error) {
	var count int64
	if err := r.db.Table("friend_requests").
		Where("sender = ? AND receiver = ?", senderName, receiverName).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) AddFriend(username string, friendName string) error {
	if err := r.db.Exec("INSERT INTO friendship (username, friend_name) VALUES (?, ?), (?, ?)",
		username, friendName, friendName, username).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFriend(username string, friendName string) error {
	if err := r.db.Exec("DELETE FROM friendship WHERE (username = ? AND friend_name = ?) OR (username = ? AND friend_name = ?)",
		username, friendName, friendName, username).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) SendFriendRequest(senderName string, receiverName string) error {
	if err := r.db.Exec("INSERT INTO friend_requests (sender, receiver, created_at) VALUES (?, ?, NOW())", senderName, receiverName).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFriendRequest(senderName string, receiverName string) error {
	if err := r.db.Exec("DELETE FROM friend_requests WHERE sender = ? AND receiver = ?", senderName, receiverName).Error; err != nil {
		return err
	}
	return nil
}
