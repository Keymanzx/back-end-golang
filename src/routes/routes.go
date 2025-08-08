package routes

import (
	_ "go-backend/docs"
	"go-backend/src/module/auth"
	"go-backend/src/module/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// global prefix for API
	api := r.Group("/api/v1")

	// Swagger docs (outside API prefix) with custom configuration
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ************ module ************
	// auth
	auth.SetupAuthRoutes(api, db)

	// user
	user.SetupUserRoutes(api, db)
}
