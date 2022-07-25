package router

import (
	"github.com/gofiber/fiber/v2"
	"rentx/models"
	"rentx/repositories"
)

func CategoriesRouter(api fiber.Router) {
	newCategoriesRepository := repositories.NewCategoriesRepository()

	api.Get("/", func(c *fiber.Ctx) error {
		categories, err := newCategoriesRepository.GetAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(categories)
	})

	api.Post("/", func(c *fiber.Ctx) error {
		body := new(models.Category)
		err := c.BodyParser(body)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
			return err
		}

		category := repositories.CreateCategoryDTO{
			Name:        body.Name,
			Description: body.Description,
		}

		checkAlreadyExists, err := newCategoriesRepository.GetByName(category.Name)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if checkAlreadyExists != nil {
			return c.Status(fiber.StatusConflict).SendString("Category already exists")
		}

		newCategory, err := newCategoriesRepository.Create(&category)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(newCategory)
	})
}
