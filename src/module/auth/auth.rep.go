package auth

import (
	"go-backend/src/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUserName(userName string) (*model.Auth, error)
	UpdateLastLogin(auth *model.Auth) error
	CreateUser(user *model.User) (*model.User,error)
	CreateAuth(auth *model.Auth) (*model.Auth,error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByUserName(userName string) (*model.Auth, error) {
	var auth model.Auth
	err := r.db.Preload("User").Where("user_name = ?", userName).First(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *authRepository) UpdateLastLogin(auth *model.Auth) error {
	return r.db.Save(auth).Error
}

func (r *authRepository) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	//find user by id
	var userGet model.User
	err = r.db.Where("id = ?", user.ID).First(&userGet).Error
	if err != nil {
		return nil, err
	}

	return &userGet, nil
}

func (r *authRepository) CreateAuth(auth *model.Auth) (*model.Auth, error) {
	err := r.db.Create(auth).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}
