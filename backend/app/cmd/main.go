package main

import (
	"github.com/gofiber/fiber/v2"
	parser "github.com/srsad/discount-service/services/parsers"
)

func main() {
  // TODO разнести по разным сервисам - все парсеры запускать кроном
  // для парсинга установить в true
	parseProduct := false
	if (parseProduct) {
    parser.Parse5Ka()
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Listen(":3000")
}
