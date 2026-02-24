package organization

import (
	"github.com/chirayusahu/queue-management-system/backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrganizationRoutes(router fiber.Router) {

	org := router.Group("/organizations", middlewares.AuthRequired)

	org.Get("/", CreateOrganization)
}
