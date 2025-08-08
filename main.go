// @title Go Backend API
// @version 1.0
// @description This is the API documentation for Go Backend
// @host localhost:8000
// @BasePath /api/v1
package main

import (
	"go-backend/src/config"
	"go-backend/src/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const signature = `
||--------------------------------------------------------------||
||██╗  ██╗███████╗██╗   ██╗███╗   ███╗ █████╗ ███╗   ██╗███████╗||
||██║ ██╔╝██╔════╝╚██╗ ██╔╝████╗ ████║██╔══██╗████╗  ██║╚══███╔╝||
||█████╔╝ █████╗   ╚████╔╝ ██╔████╔██║███████║██╔██╗ ██║  ███╔╝ ||
||██╔═██╗ ██╔══╝    ╚██╔╝  ██║╚██╔╝██║██╔══██║██║╚██╗██║ ███╔╝  ||
||██║  ██╗███████╗   ██║   ██║ ╚═╝ ██║██║  ██║██║ ╚████║███████╗||
||╚═╝  ╚═╝╚══════╝   ╚═╝   ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝||
||--------------------------------------------------------------||
`

func main() {
	config.LoadEnv()

	r := gin.Default()
	db := config.InitDB()

	routes.SetupRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	// signature
	log.Println(signature)

	// running port
	r.Run(":" + port)
}
