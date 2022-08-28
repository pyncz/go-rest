package main

import (
	"fmt"
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
	"pyncz/go-rest/utils/setup"
)

// @title Go Rest API docs
// @version 1.0
// @description API docs for my GO backend's sandbox

// @contact.name pyncz
// @contact.url http://github.com/pyncz
// @contact.email pyncz.dev@google.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	// Try to load .env vars (for dev mode)
	_ = godotenv.Load()

	logger := setup.Logger()
	db, Disconnect := setup.DB()
	defer Disconnect()

	// Create an instance of Env containing the connection pool.
	env := &models.AppEnv{
		DB:  db,
		Log: logger,
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Add routes
	{
		app.Static("/", "./public")
		app.Mount("/api/v1", api_v1.App(env))

		setup.Swagger(app.Group("/docs"))

		app.Use(favicon.New())
		app.Use(limiter.New())
		app.Use(recover.New())

		app.Use(middlewares.NotFound)
	}

	// Start server on provided port
	{
		setup.Shutdown(app, env)

		port := os.Getenv("PORT")
		if port == "" {
			port = "9090"
		}

		logger.Fatal(
			app.Listen(fmt.Sprintf(":%s", port)),
		)
	}
}
