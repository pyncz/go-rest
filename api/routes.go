package api

import (
	"pyncz/go-rest/api/apps/tasks"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	tasks.Routes(router.Group("/tasks"))
}
