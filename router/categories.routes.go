package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rentx/models"
	"time"
)

func CategoriesRouter(api fiber.Router) {
	var categories = make([]models.Category, 0)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/", func(c *fiber.Ctx) error {
		body := new(models.Category)
		err := c.BodyParser(body)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
			return err
		}

		category := models.Category{
			ID:          uuid.New().String(),
			Name:        body.Name,
			Description: body.Description,
			CreatedAt:   time.Now(),
		}

		categories = append(categories, category)

		return c.Status(fiber.StatusCreated).JSON(category)
	})
}
