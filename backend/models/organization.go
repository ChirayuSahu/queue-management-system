package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"not null"`
	Slug string    `gorm:"uniqueIndex;not null"`

	CreatedBy uuid.UUID
	Creator   User `gorm:"foreignKey:CreatedBy"`

	Members   []OrganizationMember `gorm:"foreignKey:OrganizationID"`
	Locations []Location           `gorm:"foreignKey:OrganizationID"`
	AuditLogs []AuditLog           `gorm:"foreignKey:OrganizationID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
