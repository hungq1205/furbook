package entity

type User struct {
	Username string `gorm:"primaryKey"`
	Avatar   string
	Groups   []Group `gorm:"many2many:user_groups;"`
}
