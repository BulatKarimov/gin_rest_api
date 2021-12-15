package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Event struct {
	// gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Up is executed when this migration is applied
func Up_20211215154045(txn *gorm.DB) {
	txn.CreateTable(&Event{})
}

// Down is executed when this migration is rolled back
func Down_20211215154045(txn *gorm.DB) {
	txn.DropTable(&Event{})
}
