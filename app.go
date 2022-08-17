package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	api_v1 "pyncz/go-rest/api/v1"
	"pyncz/go-rest/middlewares"
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
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	app.Static("/", "./public")
	app.Mount("/api/v1", api_v1.App(env))

	app.Use(favicon.New())
	app.Use(middlewares.NotFound)
	app.Use(limiter.New())
	app.Use(recover.New())

	// Start server on provided port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	logger.Fatal(
		app.Listen(fmt.Sprintf(":%s", port)),
	)
}
