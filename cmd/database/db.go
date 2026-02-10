package database

import (
	"log"
	"os"
	"smart-task-planner/cmd/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}
	err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})
	if err != nil {
		log.Fatal("Migration Failed")
	}
	return db
}
