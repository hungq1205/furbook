package friend

import (
	"user-service/entity"

	"gorm.io/gorm"
)

type FriendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{
		db: db,
	}
}

func (r *FriendRepository) GetFriends(username string) ([]*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Friends").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*entity.User{}, nil
		}
		return nil, err
	}
	return user.Friends, nil
}

func (r *FriendRepository) GetFriendUsernames(username string) ([]string, error) {
	friends := []string{}
	if err := r.db.
		Table("friendship").
		Where("username = ?", username).
		Pluck("friend_username", &friends).
		Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *FriendRepository) GetFriendRequests(username string) ([]*entity.User, error) {
	var reqUsers []*entity.User
	err := r.db.
		Table("friend_request fr").
		Select("u.*").
		Joins("JOIN user u ON fr.sender = u.username").
		Where("fr.receiver = ?", username).
		Find(&reqUsers).Error
	if err != nil {
		return nil, err
	}
	return reqUsers, nil
}

func (r *FriendRepository) CountFriendRequests(username string) (int64, error) {
	var count int64
	if err := r.db.Table("friend_request").
		Where("receiver = ?", username).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FriendRepository) CountFriends(username string) (int64, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FriendRepository) CheckFriendship(usernA string, userB string) (bool, error) {
	var count int64
	if err := r.db.Table("friendship").
		Where("username = ? AND friend_username = ?", usernA, userB).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *FriendRepository) CheckFriendRequest(sender string, receiver string) (bool, error) {
	var count int64
	if err := r.db.Table("friend_request").
		Where("sender = ? AND receiver = ?", sender, receiver).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *FriendRepository) AddFriend(username string, friendUsername string) error {
	if err := r.db.Exec("INSERT INTO friendship (username, friend_username) VALUES (?, ?), (?, ?)",
		username, friendUsername, friendUsername, username).Error; err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) DeleteFriend(username string, friendUsername string) error {
	if err := r.db.Exec("DELETE FROM friendship WHERE (username = ? AND friend_username = ?) OR (username = ? AND friend_username = ?)",
		username, friendUsername, friendUsername, username).Error; err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) SendFriendRequest(sender string, receiver string) error {
	if err := r.db.Exec("INSERT INTO friend_request (sender, receiver) VALUES (?, ?)", sender, receiver).Error; err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) DeleteFriendRequest(sender string, receiver string) error {
	if err := r.db.Exec("DELETE FROM friend_request WHERE sender = ? AND receiver = ?", sender, receiver).Error; err != nil {
		return err
	}
	return nil
}
