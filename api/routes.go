package api

import (
	"pyncz/go-rest/api/apps/todo"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, base string) {
	todo.TodoRoute(router, base+"/todos")
}
