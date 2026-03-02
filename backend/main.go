package main

import (
	// "context"

	"backend/config"
	"backend/database"
	// "backend/models"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	database.Connect()
	database.Migrate()

	app := fiber.New()
	routes.Setup(app, database.DB)

	app.Listen(":" + config.Get("APP_PORT"))
}
