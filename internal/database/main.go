package database

import (
	"fiscaliza/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v\n", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Report{}, &models.Recovery{}, &models.Address{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v\n", err)
	}

	return db
}
