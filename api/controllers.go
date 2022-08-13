package api

import (
	"net/http"
	"pyncz/go-rest/models"

	"github.com/gin-gonic/gin"
)

// Controllers
func Ping(env *models.AppEnv) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	}
}
