package api

import (
	"pyncz/go-rest/api/apps/tasks"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, base string) {
	tasks.Routes(router, base+"/tasks")
}
