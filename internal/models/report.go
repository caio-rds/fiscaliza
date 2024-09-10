package models

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"not null"`
	Anonymous   int    `gorm:"not null"`
	Type        string `gorm:"not null, default:'GENERIC'"`
	Description string `gorm:"not null"`
	Street      string `gorm:"not null"`
	District    string `gorm:"not null"`
	City        string `gorm:"not null"`
	State       string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Lat         string         `gorm:"not null"`
	Lon         string         `gorm:"not null"`
}

func (u *Report) TableName() string {
	return "reports"
}
