package users

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {

	users := router.Group("/users")

	users.Get("/", GetAllUsers)
	users.Post("/", CreateUser)
	users.Delete("/:id", DeleteUser)

}
