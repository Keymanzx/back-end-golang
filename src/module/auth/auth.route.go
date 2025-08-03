package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo)
	controller := NewAuthController(service)

	auth := router.Group("/auth")
	auth.POST("/register", controller.RegisterAuth)
	auth.POST("/login", controller.Login)
}
