package todo

import "github.com/gin-gonic/gin"

func TodoRoute(router *gin.Engine, base string) {
	router.GET(base+"/", Read)
	router.GET(base+"/:id", Find)
	router.POST(base+"/", Create)
}
