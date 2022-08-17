package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pyncz/go-rest/api"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
)

func main() {
	// Try to load .env vars (for dev mode)
	_ = godotenv.Load()

	// Connect db
	db, Disconnect := utils.ConnectDb()
	defer Disconnect()

	// Init logger
	logger := log.New(os.Stdout, "[go-rest] ", log.LstdFlags)

	// Create an instance of Env containing the connection pool.
	env := &models.AppEnv{
		DB:  db,
		Log: logger,
	}

	// Add routes
	router := gin.Default()
	api.Routes(router.Group("/api/v1"), env)

	// Start server on provided port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
