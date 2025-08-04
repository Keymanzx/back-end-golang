package user_test

import (
	"errors"
	"go-backend/src/model"
	"go-backend/src/module/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock UserService ---

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetProfile(userID string) (*model.User, error) {
	args := m.Called(userID)
	user, ok := args.Get(0).(*model.User)
	if !ok && args.Get(0) != nil {
		return nil, args.Error(1)
	}
	return user, args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]*model.User, error) {
	args := m.Called()
	users, ok := args.Get(0).([]*model.User)
	if !ok && args.Get(0) != nil {
		return nil, args.Error(1)
	}
	return users, args.Error(1)
}

// --- Tests ---

func TestGetProfile_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := user.NewUserController(mockService)

	uid, err := uuid.Parse("b7f85008-6207-41cb-b79a-47ef7eee204b")
	assert.NoError(t, err)

	mockUser := &model.User{
		ID:        uid,
		FirstName: "test",
		LastName:  "user",
		Email:     "test@gmail.com",
		UserType:  "ADMIN",
	}

	mockService.On("GetProfile", uid).Return(mockUser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/user/profile", nil)
	c.Set("user_id", uid)

	controller.GetProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Get Profile user")
	assert.Contains(t, w.Body.String(), "test user")

	mockService.AssertExpectations(t)
}

func TestGetProfile_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := user.NewUserController(mockService)

	mockService.On("GetProfile", "b7f85008-6207-41cb-b79a-47ef7eee204b").Return(nil, errors.New("user not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/user/profile", nil)
	c.Set("user_id", "b7f85008-6207-41cb-b79a-47ef7eee204b")

	controller.GetProfile(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")

	mockService.AssertExpectations(t)
}
