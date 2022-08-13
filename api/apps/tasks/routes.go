package tasks

import (
	"pyncz/go-rest/api/apps/tasks/items"
	"pyncz/go-rest/api/apps/tasks/tags"
	"pyncz/go-rest/models"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, env *models.AppEnv) {
	router.GET("/", Read(env))
	router.GET("/:id", Find(env))
	router.POST("/", Create(env))

	// Registed sub-domains
	tags.Routes(router.Group("/tags"), env)
	items.Routes(router.Group("/items"), env)
}
