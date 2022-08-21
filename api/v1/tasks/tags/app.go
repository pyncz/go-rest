package tags

import (
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	controller := CreateController(env)

	app.Get("/", controller.ReadPaginated)
	app.Get("/:slug", controller.FindBySlug)
	app.Post("/", controller.Create)

	return app
}
