package user

import (
	"user/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&entity.User{})
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

func (r *UserRepository) CreateUser(username string, displayName string) (*entity.User, error) {
	user := &entity.User{Username: username, DisplayName: displayName}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(username string, avatar string, bio string) (*entity.User, error) {
	if err := r.db.Model(&entity.User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{"avatar": avatar, "bio": bio}).Error; err != nil {
		return nil, err
	}
	user, err := r.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(username string) error {
	if err := r.db.Where("username = ?", username).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}
