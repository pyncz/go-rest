package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	app.Static("/", "./public")
	app.Mount("/api/v1", api.App(env))

	// Last middleware to match anything
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})

	// Start server on provided port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Fatal(
		app.Listen(fmt.Sprintf(":%s", port)),
	)
}
