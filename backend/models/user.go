package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider string

const (
	ProviderGoogle      AuthProvider = "GOOGLE"
	ProviderCredentials AuthProvider = "CREDENTIALS"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name          string
	Email         string `gorm:"uniqueIndex;not null"`
	EmailVerified bool   `gorm:"default:false"`
	PasswordHash  *string
	AuthProvider  AuthProvider `gorm:"type:varchar(20);not null"`
	Is2FAEnabled  bool         `gorm:"default:false"`

	Organizations []OrganizationMember `gorm:"foreignKey:UserID"`
	QueuesCreated []Queue              `gorm:"foreignKey:CreatedBy"`
	AuditLogs     []AuditLog           `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
