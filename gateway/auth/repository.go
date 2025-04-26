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

type Repository struct {
	db *gorm.DB
}

func NewAuthRepository(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&User{})
}

func (r *Repository) Authenticate(username, password string) (bool, error) {
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

func (r *Repository) GetUser(username string) (*User, error) {
	var user User
	err := r.db.First(&user, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(username, password string) error {
	return r.db.Create(&User{Username: username, PasswordHashed: internal.Hash(password)}).Error
}

func (r *Repository) UpdateUser(username, password string) error {
	return r.db.Save(&User{Username: username, PasswordHashed: internal.Hash(password)}).Error
}

func (r *Repository) DeleteUser(username string) error {
	return r.db.Delete(&User{}, username).Error
}
