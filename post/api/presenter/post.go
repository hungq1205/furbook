package presenter

import (
	"post/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TextPostType string

const (
	Blog  TextPostType = "blog"
	Lost  TextPostType = "lost"
	Found TextPostType = "found"
)

type Post struct {
	ID          primitive.ObjectID `json:"id"`
	Type        TextPostType       `json:"type"`
	Username    string             `json:"username"`
	DisplayName string             `json:"displayName"`
	UserAvatar  string             `json:"userAvatar"`
	Content     string             `json:"content"`
	Medias      []entity.Media     `json:"medias"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`

	Interactions []entity.Interaction `json:"interactions"`
	CommentNum   int                  `json:"commentNum"`

	// Optional: Lost Found Post
	LostAt       *time.Time       `json:"lostAt,omitempty"`
	Area         *entity.Location `json:"area,omitempty"`
	LastSeen     *entity.Location `json:"lastSeen,omitempty"`
	ContactInfo  string           `json:"contactInfo,omitempty"`
	IsResolved   bool             `json:"isResolved,omitempty"`
	Participants []string         `json:"participants,omitempty"`
}

type Comment struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}
