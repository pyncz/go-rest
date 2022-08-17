package api

import (
	"net/http"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

// Controllers
func Ping(env *models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	}
}
