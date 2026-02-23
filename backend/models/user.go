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
	ID            uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name          string       `json:"name"`
	Email         string       `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerified bool         `gorm:"default:false" json:"email_verified"`
	PasswordHash  *string      `gorm:"type:text" json:"-"`
	AuthProvider  AuthProvider `gorm:"type:varchar(20);not null" json:"auth_provider"`
	Is2FAEnabled  bool         `gorm:"default:false" json:"is_2fa_enabled"`

	Organizations []OrganizationMember `gorm:"foreignKey:UserID" json:"organizations"`
	QueuesCreated []Queue              `gorm:"foreignKey:CreatedBy" json:"queues_created"`
	AuditLogs     []AuditLog           `gorm:"foreignKey:UserID" json:"audit_logs"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
