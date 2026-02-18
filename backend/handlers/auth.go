package handlers

import (
	"strings"

	"github.com/aaksi/to-do.git/database"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(c fiber.Ctx) error {
	body := new(AuthDTO)
	err := c.Bind().JSON(body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	email := strings.TrimSpace(body.Email)
	if email == "" || !strings.Contains(email, "@") {
		return c.Status(400).JSON(fiber.Map{"error": "Некорректный email"})
	}
	if len(body.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Пароль слишком короткий"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {

		return c.Status(500).JSON(fiber.Map{"error": "Ошибка шифрования"})
	}

	user := database.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return c.Status(400).JSON(fiber.Map{"error": "Этот email уде зарегистрирован"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка сохранения"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Регистрация прошла успешно",
		"user_id": user.ID,
	})
}

func LoginHandler(c fiber.Ctx) error {
	body := new(AuthDTO)
	err := c.Bind().JSON(body)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	email := strings.TrimSpace(body.Email)

	if email == "" || !strings.Contains(email, "@") {
		return c.Status(400).JSON(fiber.Map{"error": "Некорректный формат почты"})
	}

	dbUser := database.User{}
	result := database.DB.Where("email = ?", email).First(&dbUser)
	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Неверный email или пароль"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(body.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Неверный email или пароль"})
	}

	token, err := GenerateToken(dbUser.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Не удалось создать сессию",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Успешный вход",
		"token":   token,
		"user": fiber.Map{
			"id":    dbUser.ID,
			"email": dbUser.Email,
		},
	})
}
