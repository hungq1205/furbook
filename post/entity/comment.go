package entity

import "time"

type Comment struct {
	Username  string    `bson:"username" json:"username"`
	Content   string    `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
