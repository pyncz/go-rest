package items

import (
	"pyncz/go-rest/models"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, env *models.AppEnv) {
	router.GET("/", Read(env))
	router.GET("/:id", Find(env))
	router.POST("/", Create(env))
}
