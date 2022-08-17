package tasks

import (
	"pyncz/go-rest/api/apps/tasks/items"
	"pyncz/go-rest/api/apps/tasks/tags"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	app.Get("/", Read(env))
	app.Get("/:id", Find(env))
	app.Post("/", Create(env))

	// Register sub-domains
	app.Mount("/tags", tags.App(env))
	app.Mount("/items", items.App(env))

	return app
}
