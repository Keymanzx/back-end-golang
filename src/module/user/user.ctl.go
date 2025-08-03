package user

import (
	"fmt"
	"go-backend/src/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service}
}

func (ctl *UserController) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	fmt.Println("userID", userID)
	user, err := ctl.service.GetProfile(userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// map output user
	userResponse := MapOutputUser(user)

	response.Success(c, http.StatusOK, "Get Profile user", userResponse)
}
