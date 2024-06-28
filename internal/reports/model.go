package reports

import "time"

type Report struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Anonymous bool   `gorm:"not null" default:"false"`
	Report    string `gorm:"not null"`
	Street    string `gorm:"not null"`
	District  string `gorm:"not null"`
	City      string `gorm:"not null"`
	State     string `gorm:"not null"`
	Deleted   bool   `gorm:"not null" default:"false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *Report) TableName() string {
	return "reports"
}
