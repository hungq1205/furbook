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
	Content string         `json:"content"`
	Medias  []entity.Media `json:"medias"`

	PetId    int             `json:"petId"`
	LostAt   *time.Time      `json:"lostAt"`
	Area     entity.Location `json:"area"`
	LastSeen entity.Location `json:"lastSeen"`
}

type PatchContentPayload struct {
	Content string         `json:"content"`
	Medias  []entity.Media `json:"medias"`
}

type PatchFoundPayload struct {
	Found bool `json:"found"`
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
