package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostType int

const (
	Blog PostType = iota
	LostPet
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

	// Optional: Lost Pet Post
	PetId    *int       `bson:"petId,omitempty" json:"petId,omitempty"`
	LostAt   *time.Time `bson:"lostAt,omitempty" json:"lostAt,omitempty"`
	Area     *Location  `bson:"area,omitempty" json:"area,omitempty"`
	LastSeen *Location  `bson:"lastSeen,omitempty" json:"lastSeen,omitempty"`
	Found    bool       `bson:"found,omitempty" json:"found,omitempty"`
}

type Location struct {
	Country  string `bson:"country" json:"country"`
	Province string `bson:"province" json:"province"`
	Ward     string `bson:"ward" json:"ward"`
	Street   string `bson:"street" json:"street"`
	Details  string `bson:"details" json:"details"`
}
