package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/srsad/discount-service/database"
	"github.com/srsad/discount-service/internal/entity/models"
	"github.com/srsad/discount-service/types"
)

func GetAllSource(c *fiber.Ctx) error {

	db := database.DB
	var items []models.Sources

	db.Find(&items)

	return c.JSON(fiber.Map{
		"status":  types.SUCCESS,
		"message": "",
		"data":    items,
	})
}

func GetByIdSource(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var source models.Sources
	db.First(&source, id)

	if source.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":  types.ERROR,
			"message": "Не удалось найти источник с id" + id,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  types.SUCCESS,
		"message": "",
		"data":    source,
	})
}

func CreateSource(c *fiber.Ctx) error {
	db := database.DB
	item := new(models.Sources)

	if err := c.BodyParser(item); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  types.ERROR,
			"message": "Ошибка при добавлении источника.",
			"data":    err,
		})
	}

	db.Create(&item)

	return c.Status(201).JSON(fiber.Map{
		"status":  types.SUCCESS,
		"message": "Источник успешно добавлен!",
		"data":    item,
	})
}

func UpdateSource(c *fiber.Ctx) error {
	db := database.DB

	item := new(models.Sources)

	if err := c.BodyParser(item); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  types.ERROR,
			"message": "Ошибка при обновлении источника.",
			"data":    err,
		})
	}

	var source models.Sources
	id := c.Params("id")

	db.First(&source, id)

	if source.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":  types.ERROR,
			"message": "Не удалось найти источник с id" + id,
			"data":    nil,
		})
	}

	c.BodyParser(item)
	db.Save(item)

	return c.JSON(fiber.Map{
		"status":  types.SUCCESS,
		"message": "Источник успешно обновлен!",
		"data":    item,
	})
}

func RemoveSource(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var source models.Sources
	db.First(&source, id)

	if source.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":  types.ERROR,
			"message": "Не удалось удалить источник с id" + id,
			"data":    nil,
		})
	}

	db.Delete(&source)

	return c.JSON(fiber.Map{
		"status":  types.SUCCESS,
		"message": "Источник успешно удален!",
		"data":    nil,
	})
}
