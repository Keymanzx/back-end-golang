package auth

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := gin.Default()

	// Assuming you have a login handler function
	router.POST("/login", func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Mock response for testing
		if loginRequest.UserName == "testuser" && loginRequest.Password == "testpass" {
			c.JSON(http.StatusOK, LoginResponse{Token: "mocktoken"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	t.Run("Successful Login", func(t *testing.T) {
		loginRequest := LoginRequest{
			UserName: "testuser",
			Password: "testpass",
		}
		body, _ := json.Marshal(loginRequest)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response LoginResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "mocktoken", response.Token)
	})

	t.Run("Failed Login", func(t *testing.T) {
		loginRequest := LoginRequest{
			UserName: "wronguser",
			Password: "wrongpass",
		}
		body, _ := json.Marshal(loginRequest)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
