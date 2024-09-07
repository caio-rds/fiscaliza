package database

import (
	"fiscaliza/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Report{}, &models.Recovery{}, &models.Address{})
	if err != nil {
		return nil
	}

	return db
}
