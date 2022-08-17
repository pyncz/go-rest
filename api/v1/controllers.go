package api_v1

import (
	"net/http"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

// Controllers
func Ping(*models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	}
}
