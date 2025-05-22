package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostType string

const (
	Blog  PostType = "blog"
	Lost           = "lost"
	Found          = "found"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type      PostType           `bson:"type" json:"type"`
	Username  string             `bson:"username" json:"username"`
	Content   string             `bson:"content" json:"content"`
	Medias    []Media            `bson:"medias" json:"medias"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	Interactions []Interaction `bson:"interactions" json:"interactions"`
	Comments     []Comment     `bson:"comments" json:"comments,omitempty"`

	// Optional: Lost Found Post
	LostAt       *time.Time `bson:"lostAt,omitempty" json:"lostAt,omitempty"`
	Area         *Location  `bson:"area,omitempty" json:"area,omitempty"`
	LastSeen     *Location  `bson:"lastSeen,omitempty" json:"lastSeen,omitempty"`
	ContactInfo  string     `bson:"contactInfo,omitempty" json:"contactInfo,omitempty"`
	IsResolved   bool       `bson:"isResolved,omitempty" json:"isResolved,omitempty"`
	Participants []string   `bson:"participants,omitempty" json:"participants,omitempty"`
}

type Location struct {
	Longitude float64 `bson:"longitude" json:"longitude"`
	Latitude  float64 `bson:"latitude" json:"latitude"`
}
