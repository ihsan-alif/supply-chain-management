package models

import "github.com/google/uuid"

type StockMovement struct {
	BaseModel

	ProductID     uuid.UUID  `gorm:"type:uuid;not null" json:"product_id"`
	WarehouseID   uuid.UUID  `gorm:"type:uuid;not null" json:"warehouse_id"`
	Type          string     `gorm:"type:varchar(20);not null" json:"type"`
	Qty           int        `gorm:"not null" json:"qty"`
	ReferenceType string     `gorm:"type:varchar(50)" json:"reference_type"`
	ReferenceID   *uuid.UUID `gorm:"type:uuid" json:"reference_id"`
	Notes         string     `gorm:"type:text" json:"notes"`
	CreatedBy     uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`

	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse"`
	User      User      `gorm:"foreignKey:CreatedBy" json:"user"`
}
