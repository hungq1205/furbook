package entity

type InteractionType int

const (
	Like InteractionType = iota
	Heart
)

type Interaction struct {
	Type   InteractionType `bson:"type" json:"type"`
	UserID uint            `bson:"userId" json:"user_id"`
}
