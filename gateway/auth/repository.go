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
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() (*AuthRepository, error) {
	dsn := "host=authdb user=postgres password=root dbname=auth port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &AuthRepository{
		db: db,
	}, nil
}

func (r *AuthRepository) Migrate() error {
	return r.db.AutoMigrate(&User{})
}

func (r *AuthRepository) Authenticate(username, password string) (bool, error) {
	hashedPassword := internal.Hash(password)
	var user User
	err := r.db.First(&User{}).Where("hashed_password = ?", hashedPassword).Find(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AuthRepository) GetUser(username string) (*User, error) {
	var user User
	err := r.db.First(&user, username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) CreateUser(username, password string) error {
	return r.db.Create(&User{Username: username, PasswordHashed: internal.Hash(password)}).Error
}

func (r *AuthRepository) UpdateUser(username, password string) error {
	return r.db.Save(&User{Username: username, PasswordHashed: internal.Hash(password)}).Error
}

func (r *AuthRepository) DeleteUser(username string) error {
	return r.db.Delete(&User{}, username).Error
}
