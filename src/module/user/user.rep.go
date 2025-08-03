package user

import (
	"go-backend/src/model"

	"gorm.io/gorm"
)

type UserRepository interface {
    FindAll() ([]*model.User, error)
    FindByID(userID string) (*model.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
    var users []*model.User
    err := r.db.Find(&users).Error
    return users, err
}

func (r *userRepository) FindByID(userID string) (*model.User, error) {
    var user model.User
    err := r.db.Where("id = ?", userID).First(&user).Error
    return &user, err
}
