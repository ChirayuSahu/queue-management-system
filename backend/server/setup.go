package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"
	"github.com/chirayusahu/queue-management-system/backend/routes"
)

func Setup() *fiber.App {
	cfg := config.LoadConfig()

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	database.Connect(cfg.DatabaseUrl)
	database.DB.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.Location{},
		&models.Queue{},
		&models.QueueEntry{},
		&models.TwoFactorToken{},
		&models.AuditLog{},
	)

	app.Use(logger.New())
	app.Use(recover.New())

	routes.V1Routes(app)

	app.All("*", func(c *fiber.Ctx) error {
		return common.Respond(
			c,
			fiber.StatusNotFound,
			false,
			"Cannot "+c.Method()+" "+c.OriginalURL(),
			nil,
		)
	})

	return app
}
