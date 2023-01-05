package app

import (
	"github.com/gofiber/fiber/v2"

	"github.com/srsad/discount-service/database"
	"github.com/srsad/discount-service/internal/controller/http/v1/routers"
)

func Run() {
	database.ConnectDB()

	app := fiber.New()

	routers.SetupRoutes(app)

	app.Listen(":3000")
}
