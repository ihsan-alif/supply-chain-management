package models

type Partner struct {
	BaseModel

	Code           string `gorm:"type:varchar(50);unique;not null" json:"code"`
	Type           string `gorm:"type:varchar(20);not null" json:"type"`
	CompanyName    string `gorm:"type:varchar(150);not null" json:"company_name"`
	ContactPerson  string `gorm:"type:varchar(100)" json:"contact_person"`
	Email          string `gorm:"type:varchar(100)" json:"email"`
	Phone          string `gorm:"type:varchar(20)" json:"phone"`
	City           string `gorm:"type:varchar(100)" json:"city"`
	Address        string `gorm:"type:text" json:"address"`
	Status         string `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`
	LogoURL        string `gorm:"type:text" json:"logo_url"`
}	