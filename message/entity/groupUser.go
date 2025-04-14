package entity

type GroupUser struct {
	Username string `gorm:"primaryKey"`
	GroupID  int    `gorm:"primaryKey"`
}
