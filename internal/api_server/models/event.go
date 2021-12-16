package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
