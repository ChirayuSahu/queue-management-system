package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrganizationID *uuid.UUID `gorm:"index"`
	UserID         *uuid.UUID `gorm:"index"`

	Action     string
	EntityType string
	EntityID   *uuid.UUID
	Metadata   map[string]interface{} `gorm:"type:jsonb"`

	Organization Organization `gorm:"foreignKey:OrganizationID"`
	User         User         `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
}
