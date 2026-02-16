package database

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Tasks        []Task
}

type Task struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"not null"`
	AllocatedTime int
	SpentTime     int
	IsCompleted   bool
	CreatedAt     time.Time
	UserID        uint
}
