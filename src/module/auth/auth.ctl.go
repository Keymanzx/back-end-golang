package auth

import (
	"go-backend/src/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication related operations
type AuthController struct {
	service AuthService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(service AuthService) *AuthController {
	return &AuthController{service}
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user and return token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login Request"
// @Success 201 {object} LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/login [post]
func (ctl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := ctl.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Login Success", resp)
}

// RegisterAuth godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body RegisterRequest true "Register Request"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/register [post]
func (ctl *AuthController) RegisterAuth(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := ctl.service.Register(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Registration Success", resp)
}
