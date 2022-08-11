package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pyncz/go-rest/api"
	"pyncz/go-rest/utils"
)

func main() {
	// Try to load .env vars (for dev mode)
	_ = godotenv.Load()

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

	router.Run(fmt.Sprintf(":%s", port))
}
