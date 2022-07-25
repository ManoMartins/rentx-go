package router

import (
	"github.com/gofiber/fiber/v2"
	"rentx/models"
	"rentx/repositories"
	"rentx/services"
)

func CategoriesRouter(api fiber.Router) {
	newCategoriesRepository := repositories.NewCategoriesRepository()
	newCreateCategoryService := services.NewCreateCategoryService(newCategoriesRepository)

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

		newCategory, err := newCreateCategoryService.Execute(services.Request{Name: body.Name, Description: body.Description})

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(newCategory)
	})
}
