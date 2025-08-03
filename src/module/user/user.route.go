package user

import (
	"go-backend/src/middleware/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	controller := NewUserController(service)

	userGroup := rg.Group("/user")
	userGroup.GET("/profile", jwt.AuthMiddleware(), controller.GetProfile)
}
