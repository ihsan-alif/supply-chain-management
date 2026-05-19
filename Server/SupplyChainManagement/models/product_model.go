package models

import "github.com/google/uuid"

type Product struct {
	BaseModel

	CategoryID   uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	Code         string    `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name         string    `gorm:"type:varchar(150);not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	Unit         string    `gorm:"type:varchar(30);not null" json:"unit"`
	ImageURL     string    `gorm:"type:text" json:"image_url"`
	MinimumStock int       `gorm:"default:0" json:"minimum_stock"`

	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}