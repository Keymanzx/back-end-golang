package user

import (
	"go-backend/src/model"
)

type UserService interface {
	GetProfile(userID string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetProfile(userID string) (*model.User, error) {
	return s.repo.FindByID(userID)
}

func MapOutputUser(user *model.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
