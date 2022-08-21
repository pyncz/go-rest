package tasks

import (
	"pyncz/go-rest/api/v1/tasks/items"
	"pyncz/go-rest/api/v1/tasks/tags"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	controller := CreateController(env)

	app.Get("/", controller.ReadPaginated)
	app.Get("/:id", controller.FindById)
	app.Post("/", controller.Create)

	// Register sub-domains
	app.Mount("/tags", tags.App(env))
	app.Mount("/items", items.App(env))

	return app
}
