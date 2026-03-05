package routes

import (
	"backend/handlers"
	middleware "backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func Setup(app *fiber.App, db *bun.DB) {

	userHandler := handlers.NewUserHandler(db)
	mataKuliahHandler := handlers.NewMataKuliahHandler(db)
	
	
	app.Get("/", fiber.Handler(func(c *fiber.Ctx) error {
		return c.SendString("Backend")
	}))

	
	app.Post("/users", userHandler.Register)
	app.Post("/login", userHandler.Login)

	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.Post("/logout", userHandler.Logout)
	api.Get("/users", userHandler.GetUsers)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

	api.Get("/mata-kuliah", mataKuliahHandler.GetMataKuliah)
	api.Post("/mata-kuliah", mataKuliahHandler.CreateMataKuliah)	
	api.Put("/mata-kuliah/:id", mataKuliahHandler.UpdateMataKuliah)
	api.Delete("/mata-kuliah/:id", mataKuliahHandler.DeleteMataKuliah)

}
