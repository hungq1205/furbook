package presenter

import "time"

type Message struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	GroupID   int       `json:"group_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
