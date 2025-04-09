package entity

type Group struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	IsDirect  bool
	Ownername string
}
