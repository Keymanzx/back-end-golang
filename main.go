package main

import (
	"go-backend/src/config"
	"go-backend/src/routes"

	"github.com/gin-gonic/gin"

	"os"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	db := config.InitDB()

	routes.SetupRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
