package items

import (
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	app.Get("/", Read(env))
	app.Get("/:id", Find(env))
	app.Post("/", Create(env))

	return app
}
