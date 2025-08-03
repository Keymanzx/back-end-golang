package routes

import (
	"go-backend/src/module/auth"
	"go-backend/src/module/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	//global prefix
	api := r.Group("/api/v1")

	// ************ module ************

	// auth
	auth.SetupAuthRoutes(api, db)

	// user
	user.SetupUserRoutes(api, db)
}
