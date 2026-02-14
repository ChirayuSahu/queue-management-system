package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;column:user_id"`
	Name     string    `json:"name"`
	Username string    `json:"username" gorm:"uniqueIndex;not null"`
	Password string    `json:"-"`
	Email    string    `json:"email" gorm:"uniqueIndex;not null"`
	IsAdmin  string    `json:"is_admin" gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
