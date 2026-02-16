package main

import (
	"fmt"

	"github.com/aaksi/to-do.git/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	fmt.Println("Application started successful")

	//newUser := database.User{
	//	Email:        "stas2@mail.ru",
	//	PasswordHash: "password_hash",
	//}
	//
	//result := database.DB.Create(&newUser)
	//
	//if result.Error != nil {
	//	fmt.Println("Ошибка при создании пользователя:", result.Error)
	//} else {
	//	fmt.Printf("Пользователь создан! ID: %d, Статус: %v\n", newUser.ID, newUser.Email)
	//}
	//
	//newTask := database.Task{
	//	Title:         "Изучить GORM",
	//	AllocatedTime: 60,
	//	IsCompleted:   true,
	//	UserID:        newUser.ID,
	//}
	//
	//result = database.DB.Create(&newTask)
	//
	//if result.Error != nil {
	//	fmt.Println("Ошибка при создании задачи:", result.Error)
	//} else {
	//	fmt.Printf("Задача создана! ID: %d, Статус: %v\n", newTask.ID, newTask.IsCompleted)
	//}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("fiber")
	})
	fmt.Println("Сервер v3 запущен на порту 3000")
	app.Listen(":3000")
}
