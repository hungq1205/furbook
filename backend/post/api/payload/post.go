package payload

import (
	"post/entity"
	"time"
)

type UsersPostsRequest struct {
	Usernames []string `json:"usernames"`
}

type CreateBlogPostPayload struct {
	Content string         `json:"content"`
	Medias  []entity.Media `json:"medias"`
}

type CreateLostPetPostPayload struct {
	Type    entity.PostType `json:"type"`
	Content string          `json:"content"`
	Medias  []entity.Media  `json:"medias"`

	ContactInfo string          `json:"contactInfo"`
	LostAt      *time.Time      `json:"lostAt"`
	Area        entity.Location `json:"area"`
	LastSeen    entity.Location `json:"lastSeen"`
}

type PatchContentPayload struct {
	Content string         `json:"content"`
	Medias  []entity.Media `json:"medias"`
}

type PatchLostFoundStatus struct {
	IsResolved bool `json:"isResolved"`
}

type DeletePostPayload struct {
	PostID string `json:"postId"`
}

type CreateCommentPayload struct {
	Content string `json:"content"`
}

type UpsertInteractionPayload struct {
	Type entity.InteractionType `json:"type"`
}
