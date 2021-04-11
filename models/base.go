package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
