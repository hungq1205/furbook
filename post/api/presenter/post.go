package presenter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/entity"
	"time"
)

type textPostType string

const (
	Blog  textPostType = "blog"
	Lost  textPostType = "lost"
	Found textPostType = "found"
)

type Post struct {
	ID         primitive.ObjectID `json:"id"`
	Type       textPostType       `json:"type"`
	UserID     uint               `json:"userId"`
	Username   string             `json:"username"`
	UserAvatar string             `json:"userAvatar"`
	Content    string             `json:"content"`
	Medias     []entity.Media     `json:"medias"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`

	Interactions []entity.Interaction `json:"interactions"`
	CommentNum   int                  `json:"commentNum"`

	// Optional: Lost Found Post
	LostAt       *time.Time       `json:"lostAt,omitempty"`
	Area         *entity.Location `json:"area,omitempty"`
	LastSeen     *entity.Location `json:"lastSeen,omitempty"`
	ContactInfo  string           `json:"contactInfo,omitempty"`
	IsResolved   bool             `json:"isResolved,omitempty"`
	Participants []uint           `json:"participants,omitempty"`
}
