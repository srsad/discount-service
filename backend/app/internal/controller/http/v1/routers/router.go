package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/srsad/discount-service/internal/controller/http/v1/handler"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), cors.New())
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api")
	})

	// // Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)

	// // User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// Source
	source := api.Group("/source")
	source.Get("/", handler.GetAllSource)
	source.Get("/:id", handler.GetByIdSource)
	source.Post("/", handler.CreateSource)
	source.Put("/:id", handler.UpdateSource)
	source.Delete("/:id", handler.RemoveSource)
}
