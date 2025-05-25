package entity

import "time"

type Message struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	GroupID   int
	Group     Group  `gorm:"foreignKey:GroupID"`
	Content   string `gorm:"size:1024"`
	CreatedAt time.Time
}
