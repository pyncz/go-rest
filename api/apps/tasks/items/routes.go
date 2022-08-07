package items

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/", Read)
	router.GET("/:id", Find)
	router.POST("/", Create)
}
