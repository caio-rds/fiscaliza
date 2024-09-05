package database

import "gorm.io/gorm"

type DbStruct struct {
	*gorm.DB
}

func NewDb(db *gorm.DB) *DbStruct {
	value := DbStruct{
		db,
	}
	return &value
}
