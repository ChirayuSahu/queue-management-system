package models

import (
	"time"

	"github.com/google/uuid"
)

type EntryStatus string

const (
	EntryWaiting   EntryStatus = "WAITING"
	EntryCalled    EntryStatus = "CALLED"
	EntryServed    EntryStatus = "SERVED"
	EntryCancelled EntryStatus = "CANCELLED"
	EntryNoShow    EntryStatus = "NO_SHOW"
)

type QueueEntry struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	QueueID     uuid.UUID `gorm:"not null;index"`
	TokenNumber int       `gorm:"not null"`

	Name   string `gorm:"not null"`
	Phone  *string
	Email  *string
	Status EntryStatus `gorm:"type:varchar(20);default:'WAITING'"`

	Queue Queue `gorm:"foreignKey:QueueID;constraint:OnDelete:CASCADE"`

	JoinedAt    time.Time
	CalledAt    *time.Time
	ServedAt    *time.Time
	CancelledAt *time.Time
}
