package models

import "gorm.io/gorm"

type Recovery struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Username  string
	Code      string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
