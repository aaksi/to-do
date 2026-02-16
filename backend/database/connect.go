package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(errors.New("connection failed"))
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pass, dbname, port, sslmode)

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{}, &Task{})
	if err != nil {
		panic(err)
	}
	DB = db

	fmt.Println("Database connection and migration successful")
}
