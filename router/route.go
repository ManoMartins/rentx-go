package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	app.Route("/categories", CategoriesRouter, "categories")
}
