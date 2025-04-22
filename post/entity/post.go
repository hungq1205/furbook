package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostType int

const (
	Blog PostType = iota
	Lost
	Found
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type      PostType           `bson:"type" json:"type"`
	UserID    uint               `bson:"userId" json:"userId"`
	Content   string             `bson:"content" json:"content"`
	Medias    []Media            `bson:"medias" json:"medias"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	Interactions []Interaction `bson:"interactions" json:"interactions"`
	Comments     []Comment     `bson:"comments" json:"comments,omitempty"`

	// Optional: Lost Found Post
	LostAt      *time.Time `bson:"lostAt,omitempty" json:"lostAt,omitempty"`
	Area        *Location  `bson:"area,omitempty" json:"area,omitempty"`
	LastSeen    *Location  `bson:"lastSeen,omitempty" json:"lastSeen,omitempty"`
	ContactInfo string     `bson:"contactInfo,omitempty" json:"contactInfo,omitempty"`
	IsResolved  bool       `bson:"found,omitempty" json:"found,omitempty"`
	Partipants  []uint     `bson:"participants,omitempty" json:"participants,omitempty"`
}

type Location struct {
	Country  string `bson:"country" json:"country"`
	Province string `bson:"province" json:"province"`
	Ward     string `bson:"ward" json:"ward"`
	Street   string `bson:"street" json:"street"`
	Details  string `bson:"details" json:"details"`
}
