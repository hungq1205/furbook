package entity

type MediaType string

const (
	Image MediaType = "image"
	Video           = "video"
)

type Media struct {
	Type MediaType `bson:"type" json:"type"`
	URL  string    `bson:"url" json:"url"`
}
