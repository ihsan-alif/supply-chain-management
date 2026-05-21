package models

import "github.com/google/uuid"

type ProductStock struct {
	BaseModel

	ProductID   uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null" json:"warehouse_id"`
	Stock       int       `gorm:"default:0" json:"stock"`

	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse"`
}