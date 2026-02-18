package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c fiber.Ctx) error {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Ошибка аутентификации"})
	}

	authHeader := c.Get("Authorization")

	if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
		return c.Status(401).JSON(fiber.Map{"error": "Ошибка аутентификации"})
	}

	tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
	//t, err := jwt.Parse(secret, func())
	return c.Next()
}
