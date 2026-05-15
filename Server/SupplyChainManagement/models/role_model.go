package models

type Role struct {
	BaseModel
	
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}
