package models

import "github.com/google/uuid"

type User struct {
	BaseModel

	RoleID   uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	FullName string    `gorm:"type:varchar(100);not null" json:"full_name"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string    `gorm:"type:text;not null" json:"-"`
	Avatar   string    `gorm:"type:text" json:"avatar"`
	Phone    string    `gorm:"type:varchar(20)" json:"phone"`
	IsActive bool      `gorm:"default:true" json:"is_active"`

	Role Role `gorm:"foreignKey:RoleID"`
}
