package main

import (
	"log"

	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/server"
)

func main() {
	app := server.Setup()

	cfg := config.LoadConfig()
	port := cfg.Port
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
