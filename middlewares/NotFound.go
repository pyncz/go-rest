package middlewares

import "github.com/gofiber/fiber/v2"

// Middleware to match any route
func NotFound(ctx *fiber.Ctx) error {
	return ctx.SendStatus(404)
}
