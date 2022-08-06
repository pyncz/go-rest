package tasks

import (
	"pyncz/go-rest/api/apps/tasks/items"
	"pyncz/go-rest/api/apps/tasks/tags"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, base string) {
	router.GET(base+"/", Read)
	router.GET(base+"/:id", Find)
	router.POST(base+"/", Create)

	// Registed sub-domains
	tags.Routes(router, base+"/tags")
	items.Routes(router, base+"/items")
}
