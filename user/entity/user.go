package entity

import "time"

type User struct {
	ID       uint `gorm:"PrimaryKey"`
	Username string
	Avatar   string
	Friends  []*User `gorm:"many2many:friendship;joinForeignKey:ID;JoinReferences:ID"`
}

type FriendRequest struct {
	Sender    string `gorm:"PrimaryKey"`
	Receiver  string `gorm:"PrimaryKey"`
	CreatedAt time.Time
}
