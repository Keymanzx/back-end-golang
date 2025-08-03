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
