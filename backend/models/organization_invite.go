package models

import (
	"time"

	"github.com/google/uuid"
)

type InviteStatus string

const (
	InvitePending  InviteStatus = "PENDING"
	InviteAccepted InviteStatus = "ACCEPTED"
	InviteExpired  InviteStatus = "EXPIRED"
	InviteRevoked  InviteStatus = "REVOKED"
)

type OrganizationInvite struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrganizationID uuid.UUID      `gorm:"not null;index"`
	Email          string         `gorm:"not null"`
	Role           MembershipRole `gorm:"type:varchar(20);not null"`
	Token          string         `gorm:"uniqueIndex;not null"`
	Status         InviteStatus   `gorm:"type:varchar(20);default:'PENDING'"`
	ExpiresAt      time.Time
	CreatedBy      *uuid.UUID

	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
}
