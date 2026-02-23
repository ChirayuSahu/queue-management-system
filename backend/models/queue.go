package models

import (
	"time"

	"github.com/google/uuid"
)

type QueueStatus string

const (
	QueueActive QueueStatus = "ACTIVE"
	QueuePaused QueueStatus = "PAUSED"
	QueueClosed QueueStatus = "CLOSED"
)

type Queue struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LocationID uuid.UUID  `gorm:"not null;index"`
	CreatedBy  *uuid.UUID `gorm:"index"`

	Name        string `gorm:"not null"`
	Description *string
	Status      QueueStatus `gorm:"type:varchar(20);default:'ACTIVE'"`
	IsPublic    bool        `gorm:"default:true"`
	PublicSlug  *string     `gorm:"uniqueIndex"`

	Location Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE"`
	Creator  User     `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL"`

	Entries []QueueEntry `gorm:"foreignKey:QueueID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
