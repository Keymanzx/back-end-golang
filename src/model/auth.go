package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserName  string    `gorm:"user_name" gorm:"unique"`
	Password  string    `gorm:"password"`
	LastLogin time.Time `gorm:"last_login"`

	// foreign key
	UserID uuid.UUID `gorm:"user_id"`
	User   *User     `gorm:"user"`
}

func (u *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
