package main

import (
	"fmt"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}
