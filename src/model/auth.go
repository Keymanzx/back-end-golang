package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserName  string    `json:"user_name" gorm:"unique"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`

	// foreign key
	UserID uuid.UUID `json:"user_id"`
}

func (u *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
