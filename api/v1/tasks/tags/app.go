package tags

import (
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	app.Get("/", Read(env))
	app.Get("/:slug", Find(env))
	app.Post("/", Create(env))

	return app
}
