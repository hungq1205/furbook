package auth

import (
	"errors"
	"gateway/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Username       string `gorm:"primary_key"`
	PasswordHashed string `gorm:"password_hashed"`
	Salt           string `gorm:"salt"`
}

type Repository struct {
	db *gorm.DB
}

func NewAuthRepository(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Authenticate(username, password string) (bool, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return internal.CompareHashAndPassword(user.PasswordHashed, password, user.Salt), nil
}

func (r *Repository) GetUser(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(username, password string) error {
	salt := internal.GenerateSalt()
	hashedPassword := internal.Hash(password, salt)
	return r.db.Create(&User{Username: username, PasswordHashed: hashedPassword, Salt: salt}).Error
}

func (r *Repository) UpdateUser(username, password string) error {
	salt := internal.GenerateSalt()
	hashedPassword := internal.Hash(password, salt)
	return r.db.Save(&User{Username: username, PasswordHashed: hashedPassword, Salt: salt}).Error
}

func (r *Repository) DeleteUser(username string) error {
	return r.db.Delete(&User{}, username).Error
}
