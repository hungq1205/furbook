package entity

import "time"

type Notification struct {
	ID        int       `gorm:"PrimaryKey;autoIncrement" json:"id"`
	Username  string    `json:"username"`
	Icon      string    `json:"icon"`
	Desc      string    `json:"desc"`
	Link      string    `json:"link"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}
