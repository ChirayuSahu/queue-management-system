package main

import (
	"log"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App

func setupApp() *fiber.App {
	cfg := config.LoadConfig()

	app = fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	database.Connect(cfg.DatabaseUrl)

	// database.DB.AutoMigrate(
	// 	&models.User{},
	// )

	app.Use(logger.New())
	app.Use(recover.New())
	routes.V1Routes(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return common.Respond(c, fiber.StatusOK, true, "API is healthy", nil)
	})

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

func main() {
	cfg := config.LoadConfig()

	app := setupApp()

	log.Println("Server running on http://localhost:" + cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}

func Handler(c *fiber.Ctx) error {
	if app == nil {
		setupApp()
	}
	app.Handler()(c.Context())
	return nil
}
