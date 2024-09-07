package reports

import "gorm.io/gorm"

type StructRep struct {
	*gorm.DB
}

func NewDb(db *gorm.DB) *StructRep {
	value := StructRep{
		db,
	}
	return &value
}
