package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// create uuid before insert data
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}