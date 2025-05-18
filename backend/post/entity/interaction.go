package entity

type InteractionType string

const (
	Like  InteractionType = "like"
	Heart                 = "heart"
)

type Interaction struct {
	Type     InteractionType `bson:"type" json:"type"`
	Username string          `bson:"username" json:"username"`
}
