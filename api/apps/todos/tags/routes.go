package tags

import "github.com/gin-gonic/gin"

func TagsRoute(router *gin.Engine, base string) {
	router.GET(base+"/", Read)
	router.GET(base+"/:slug", Find)
	router.POST(base+"/", Create)
}
