package organization

import (
	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	return common.Respond(c, fiber.StatusOK, true, "Route Working", nil)
}
