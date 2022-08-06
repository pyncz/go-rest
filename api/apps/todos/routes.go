package todos

import (
	"pyncz/go-rest/api/apps/todos/tags"

	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine, base string) {
	router.GET(base+"/", Read)
	router.GET(base+"/:id", Find)
	router.POST(base+"/", Create)

	// Registed sub-domains
	tags.TagsRoute(router, base+"/tags")
}
