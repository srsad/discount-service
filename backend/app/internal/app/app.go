package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srsad/discount-service/internal/controller/http/v1/routers"
)

func Run() {
	app := fiber.New()

	routers.SetupRoutes(app)

	app.Listen(":3000")
}
