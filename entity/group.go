package entity

type Group struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	IsDirect bool
	Members  []User `gorm:"many2many:user_groups;"`
}
