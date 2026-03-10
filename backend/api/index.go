package api

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/chirayusahu/queue-management-system/backend/server"
)

var (
	app  *fiber.App
	once sync.Once
)

func buildApp() *fiber.App {
	once.Do(func() {
		app = server.Setup()
	})
	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	adaptor.FiberApp(buildApp())(w, r)
}
