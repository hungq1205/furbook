package entity

import "time"

type Comment struct {
	UserID    uint      `bson:"userId" json:"user_id"`
	Content   string    `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
