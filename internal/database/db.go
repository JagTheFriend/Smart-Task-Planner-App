package database

import (
	"log/slog"
	"os"
	"smart-task-planner/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}
	err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})
	if err != nil {
		slog.Error("Database: Failed Migration")
	}
	return db
}
