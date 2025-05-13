package entity

type MediaType int

const (
	Image MediaType = iota
	Video
)

type Media struct {
	Type MediaType `bson:"type" json:"type"`
	URL  string    `bson:"url" json:"url"`
}
