package entity

import "time"

type User struct {
	Username    string `gorm:"PrimaryKey"`
	DisplayName string
	Bio         string
	Avatar      string
	Friends     []*User `gorm:"many2many:friendship;joinForeignKey:Username;JoinReferences:FriendName"`
}

type FriendRequest struct {
	Sender    string `gorm:"PrimaryKey"`
	Receiver  string `gorm:"PrimaryKey"`
	CreatedAt time.Time
}
