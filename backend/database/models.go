package database

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Tasks        []Task `gorm:"foreignKey:UserID"`
}

type Task struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"not null"`
	IsTimed       bool   `gorm:"default:false"`
	AllocatedTime int
	SpentTime     int

	IsCompleted bool

	StartDate *time.Time
	FinishAt  *time.Time
	CreatedAt time.Time

	UserID uint

	ParentID *uint
	SubTasks []Task `gorm:"foreignKey:ParentID"`
}
