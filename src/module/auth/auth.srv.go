package auth

import (
	"errors"
	"go-backend/src/middleware/jwt"
	"go-backend/src/model"
	"go-backend/src/utils/hash"
	"time"
)

type AuthService interface {
	Login(req LoginRequest) (LoginResponse, error)
	Register(req RegisterRequest) (RegisterResponse, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req LoginRequest) (LoginResponse, error) {
	auth, err := s.repo.FindByUserName(req.UserName)
	if err != nil {
		return LoginResponse{}, errors.New("invalid username or password")
	}

	// compare password
	err = hash.ComparePassword(auth.Password, req.Password)
	if err != nil {
		return LoginResponse{}, errors.New("invalid username or password")
	}

	token, err := jwt.GenerateToken(auth.UserID.String())
	if err != nil {
		return LoginResponse{}, errors.New("failed to generate token")
	}

	auth.LastLogin = time.Now()
	s.repo.UpdateLastLogin(auth)

	return LoginResponse{Token: token}, nil
}

func (s *authService) Register(req RegisterRequest) (RegisterResponse, error) {
	user := &model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserType:  "ADMIN",
	}

	user, err := s.repo.CreateUser(user)
	if err != nil {
		return RegisterResponse{}, errors.New("failed to create user")
	}

	// hash password
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return RegisterResponse{}, errors.New("failed to hash password")
	}

	auth := &model.Auth{
		UserName: req.UserName,
		Password: hashedPassword,
		UserID:   user.ID,
	}

	_, err = s.repo.CreateAuth(auth)
	if err != nil {
		return RegisterResponse{}, errors.New("failed to create auth")
	}

	token, err := jwt.GenerateToken(user.ID.String())
	if err != nil {
		return RegisterResponse{}, errors.New("failed to generate token")
	}

	return RegisterResponse{Token: token}, nil
}
