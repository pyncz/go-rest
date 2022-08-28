package setup

import (
	_ "pyncz/go-rest/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// Setup swagger docs. See details:
// https://github.com/gofiber/swagger#canonical-example
func Swagger(app fiber.Router) {
	app.Get("/*", swagger.New(swagger.Config{
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "list",
	}))
}
