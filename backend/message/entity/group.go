package entity

type Group struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	IsDirect  bool
	OwnerName string
}
