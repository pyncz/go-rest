package api

import (
	"pyncz/go-rest/api/apps/todos"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, base string) {
	todos.TodoRoute(router, base+"/todos")
}
