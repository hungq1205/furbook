package user

import (
	"user-service/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUser(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers(usernames []string) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Where("username IN ?", usernames).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("Username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) CreateUser(username string, avatar string) (*entity.User, error) {
	user := &entity.User{Username: username, Avatar: avatar}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(username string, avatar string) (*entity.User, error) {
	if err := r.db.Model(&entity.User{}).
		Where("username = ?", username).
		Update("avatar", avatar).Error; err != nil {
		return nil, err
	}
	user, err := r.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(username string) error {
	if err := r.db.Delete(&entity.User{}, username).Error; err != nil {
		return err
	}
	return nil
}
