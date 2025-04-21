package entity

type InteractionType int

const (
	Like InteractionType = iota
	Heart
)

type Interaction struct {
	Type   InteractionType `bson:"type" json:"type"`
	userId string          `bson:"user_id" json:"user_id"`
}
