package tags

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine, base string) {
	router.GET(base+"/", Read)
	router.GET(base+"/:slug", Find)
	router.POST(base+"/", Create)
}
