package entity

import "time"

type Message struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	User      User `gorm:"foreignKey:Username"`
	GroupID   int
	Group     Group  `gorm:"foreignKey:GroupID"`
	Content   string `gorm:"size:255"`
	CreatedAt time.Time
}
