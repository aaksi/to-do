package main

import (
	"fmt"

	"github.com/aaksi/to-do.git/database"
	"github.com/aaksi/to-do.git/handlers"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	//newUser := database.User{
	//	Email:        "test2@mail.ru",
	//	PasswordHash: "pass",
	//}
	//
	//result := database.DB.Create(&newUser)
	//
	//if result.Error != nil {
	//	log.Fatal("ошибка создания пользователя", result.Error)
	//}
	//
	//fmt.Println("Пользователь создан, id: ", newUser.ID)
	//
	//now := time.Now()
	//
	//newTask := database.Task{
	//	Title:     "task2 test",
	//	StartDate: &now,
	//	UserID:    newUser.ID,
	//}
	//
	//result = database.DB.Create(&newTask)
	//
	//if result.Error != nil {
	//	log.Fatal("ошибка создания задачи", result.Error)
	//}
	//
	//fmt.Println("Задача создана, id: ", newTask.ID)
	//
	//newSubTask := database.Task{
	//	Title:     "subtask2 test",
	//	StartDate: &now,
	//	ParentID:  &newTask.ID,
	//	UserID:    newUser.ID,
	//}
	//result = database.DB.Create(&newSubTask)
	//
	//if result.Error != nil {
	//	log.Fatal("ошибка создания подзадачи", result.Error)
	//}
	//
	//fmt.Println("подзадача создана, id: ", newSubTask.ID)
	//newSubTask = database.Task{
	//	Title:     "subtask3 test",
	//	StartDate: &now,
	//	ParentID:  &newTask.ID,
	//	UserID:    newUser.ID,
	//}
	//result = database.DB.Create(&newSubTask)
	//
	//if result.Error != nil {
	//	log.Fatal("ошибка создания подзадачи", result.Error)
	//}
	//
	//fmt.Println("подзадача создана, id: ", newSubTask.ID)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("fiber")
	})
	app.Post("/register", handlers.RegisterHandler)
	app.Get("/login", handlers.LoginHandler)
	fmt.Println("Сервер v3 запущен на порту 3000")
	app.Listen(":3000")
}
