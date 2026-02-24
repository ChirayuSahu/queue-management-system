package middlewares

import (
	"fmt"
	"strings"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Missing Authorization header", nil)
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Invalid Authorization format", nil)
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	cfg := config.LoadConfig()
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Invalid or expired token", nil)
	}

	if claims.ID == "" {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Invalid token payload", nil)
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", claims.ID).Error; err != nil {
		return common.Respond(c, fiber.StatusUnauthorized, false, "User not found", nil)
	}

	c.Locals("user", user)
	return c.Next()
}
