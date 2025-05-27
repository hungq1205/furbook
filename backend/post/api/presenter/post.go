package presenter

import (
	"post/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `json:"id"`
	Type        entity.PostType    `json:"type"`
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
	LostAt       *time.Time `json:"lostAt,omitempty"`
	Area         *Location  `json:"area,omitempty"`
	LastSeen     *Location  `json:"lastSeen,omitempty"`
	ContactInfo  string     `json:"contactInfo,omitempty"`
	IsResolved   bool       `json:"isResolved,omitempty"`
	Participants []string   `json:"participants,omitempty"`
}

type Location struct {
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type Comment struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

func LocationEntityToPresenter(loc *entity.Location) *Location {
	if loc == nil {
		return nil
	}
	return &Location{
		Address: loc.Address,
		Lat:     loc.Location.Coordinates[1],
		Lng:     loc.Location.Coordinates[0],
	}
}
