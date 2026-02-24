package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", RegisterUser)
	auth.Post("/login", LoginUser)
}
