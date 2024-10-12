package models

import (
	"gorm.io/gorm"
	"time"
)

type Recovery struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Code      string
	Type      string
	Target    string
	MessageId string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Recovery) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ExpiresAt.IsZero() {
		r.ExpiresAt = time.Now().Add(30 * time.Minute)
	}
	return
}
