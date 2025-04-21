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

func (r *UserRepository) GetUser(userId uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers(userIds []uint) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Where("id IN ?", userIds).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(username string, avatar string) (*entity.User, error) {
	user := &entity.User{Username: username, Avatar: avatar}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(userId uint, avatar string) (*entity.User, error) {
	if err := r.db.Model(&entity.User{}).
		Where("id = ?", userId).
		Update("avatar", avatar).Error; err != nil {
		return nil, err
	}
	user, err := r.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(userId uint) error {
	if err := r.db.Delete(&entity.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}
