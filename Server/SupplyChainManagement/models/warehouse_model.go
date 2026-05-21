package models

type Warehouse struct {
	BaseModel

	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Address string `gorm:"type:text" json:"address"`
}