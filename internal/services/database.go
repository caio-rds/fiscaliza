package services

import (
	"fiscaliza/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := "root:rc321@tcp(127.0.0.1:3306)/community_voice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Report{}, &models.Recovery{})
	if err != nil {
		return nil
	}

	return db
}

//type dbInterface struct {
//	*gorm.DB
//}
//
//func DB(db *gorm.DB) *dbInterface {
//	value := dbInterface{
//		db,
//	}
//	return &value
//}
