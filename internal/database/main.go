package database

import (
	"fiscaliza/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fiscaliza.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Report{}, &models.Recovery{}, &models.Address{})
	if err != nil {
		log.Fatalln("failed to migrate database: %v", err)
	}

	return db
}
