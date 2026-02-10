package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Tasks     []Task
	CreatedAt time.Time
}
