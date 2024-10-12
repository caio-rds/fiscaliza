package models

type Address struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"not null"`
	Street     string
	Compliment *string
	District   string
	City       string
	State      string
	Name       string
	Default    bool
	Lat        string
	Lon        string
}

func (a *Address) TableName() string {
	return "user_addresses"
}
