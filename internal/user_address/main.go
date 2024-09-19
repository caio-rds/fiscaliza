package user_address

import "gorm.io/gorm"

type Struct struct {
	*gorm.DB
}

func NewDb(db *gorm.DB) *Struct {
	value := Struct{
		db,
	}
	return &value
}
