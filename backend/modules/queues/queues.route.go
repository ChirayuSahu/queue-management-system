package queues

import (
	"github.com/gofiber/fiber/v2"
)

func QueueRoutes(router fiber.Router) {

	queues := router.Group("/queues")

	queues.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Get all queues")
	})

}
