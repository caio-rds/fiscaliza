package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:rc321@tcp(127.0.0.1:3306)/community_voice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//err = db.AutoMigrate(CreateDB{})
	//if err != nil {
	//	return
	//}

	return db
}
