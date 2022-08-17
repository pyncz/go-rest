package api

import (
	"pyncz/go-rest/api/apps/tasks"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func App(env *models.AppEnv) *fiber.App {
	app := fiber.New()

	// System routes
	app.Get("/ping", Ping(env))

	// Register sub-domains
	app.Mount("/tasks", tasks.App(env))

	return app
}
