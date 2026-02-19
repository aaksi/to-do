package middleware

import (
	"fmt"
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
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Ошибка аутентификации"})
	}

	if !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Ошибка структуры токена"})
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "ID пользователя не найден в токене"})
	}

	c.Locals("user_id", uint(userID))
	return c.Next()
}
