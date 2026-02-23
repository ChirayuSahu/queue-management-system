package middlewares

import (
	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"
	"github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	if userID == nil {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Unauthorized", nil)
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Unauthorized", nil)
	}

	c.Locals("user", user)
	return c.Next()
}
