package main

import (
	"github.com/gin-gonic/gin"

	"pyncz/go-rest/api"
)

func main() {
	router := gin.Default()

	api.RegisterRoutes(router, "/api/v1")

	router.Run("localhost:9090")
}
