package tags

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/", Read)
	router.GET("/:slug", Find)
	router.POST("/", Create)
}
