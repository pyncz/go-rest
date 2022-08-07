package tasks

import (
	"pyncz/go-rest/api/apps/tasks/items"
	"pyncz/go-rest/api/apps/tasks/tags"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/", Read)
	router.GET("/:id", Find)
	router.POST("/", Create)

	// Registed sub-domains
	tags.Routes(router.Group("/tags"))
	items.Routes(router.Group("/items"))
}
