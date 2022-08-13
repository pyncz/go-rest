package tags

import (
	"pyncz/go-rest/models"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, env *models.AppEnv) {
	router.GET("/", Read(env))
	router.GET("/:slug", Find(env))
	router.POST("/", Create(env))
}
