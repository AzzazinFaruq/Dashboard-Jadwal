package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func Setup(app *fiber.App, db *bun.DB) {

	userHandler := handlers.NewUserHandler(db)
	mataKuliahHandler := handlers.NewMataKuliahHandler(db)
	
	
	app.Get("/", fiber.Handler(func(c *fiber.Ctx) error {
		return c.SendString("Backend")
	}))

	api := app.Group("/api")

	api.Get("/users", userHandler.GetUsers)
	api.Post("/users", userHandler.CreateUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

	api.Get("/mata-kuliah", mataKuliahHandler.GetMataKuliah)
	api.Post("/mata-kuliah", mataKuliahHandler.CreateMataKuliah)


}
