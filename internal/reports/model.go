package reports

import "time"

type Report struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"not null"`
	Anonymous int       `gorm:"not null"`
	Report    string    `gorm:"not null"`
	Street    string    `gorm:"not null"`
	District  string    `gorm:"not null"`
	City      string    `gorm:"not null"`
	State     string    `gorm:"not null"`
	Deleted   int       `gorm:"not null"`
	CreatedAt time.Time `column:"created_at" `
	UpdatedAt time.Time `column:"updated_at"`
	DeletedAt time.Time `column:"deleted_at" default:"0000-00-00 00:00:00"`
}

func (u *Report) TableName() string {
	return "reports"
}
