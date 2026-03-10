package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrganizationID uuid.UUID `gorm:"not null;index"`

	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE"`

	Name    string `gorm:"not null"`
	Address *string

	Queues []Queue `gorm:"foreignKey:LocationID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
