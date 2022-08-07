package main

import (
	"github.com/gin-gonic/gin"

	"pyncz/go-rest/api"
	"pyncz/go-rest/utils"
)

func main() {
	// Connect db
	Disconnect := utils.ConnectDb()
	defer Disconnect()

	// Add routes
	router := gin.Default()
	api.Routes(router.Group("/api/v1"))

	router.Run("localhost:9090")
}
