package api

import (
	"pyncz/go-rest/api/apps/tasks"
	"pyncz/go-rest/models"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, env *models.AppEnv) {
	// System routes
	router.GET("/ping", Ping(env))

	// Apps
	tasks.Routes(router.Group("/tasks"), env)
}
