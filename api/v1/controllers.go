package api_v1

import (
	"net/http"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

/*
 * Controllers
 */

// Health-check.
//
// @Summary health check
// @Description Responses with 200 OK if service is available.
// @Tags Health
// @Success 200
// @Failure      503  {object}  utils.HttpError
// @Router /ping [get]
func Ping(*models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	}
}
