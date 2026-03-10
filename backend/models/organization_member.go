package models

import (
	"time"

	"github.com/google/uuid"
)

type MembershipRole string

const (
	RoleOwner      MembershipRole = "OWNER"
	RoleAdmin      MembershipRole = "ADMIN"
	RoleSupervisor MembershipRole = "SUPERVISOR"
)

type OrganizationMember struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrganizationID uuid.UUID      `gorm:"not null;index"`
	UserID         uuid.UUID      `gorm:"not null;index"`
	Role           MembershipRole `gorm:"type:varchar(20);not null"`

	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE"`
	User         User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time

	_ struct{} `gorm:"uniqueIndex:idx_org_user,composite:org_user"`
}
