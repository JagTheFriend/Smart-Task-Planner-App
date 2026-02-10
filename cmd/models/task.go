package models

import "time"

type Task struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	CategoryID  uint
	Category    Category
	Deadline    time.Time
	Completed   bool
	UserID      uint
	CreatedAt   time.Time
}
