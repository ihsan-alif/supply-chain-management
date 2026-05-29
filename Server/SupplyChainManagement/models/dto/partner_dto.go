package dto

type PartnerRequest struct {
	Code          string `json:"code" binding:"required"`
	Type          string `json:"type" binding:"required,oneof=SUPPLIER CUSTOMER DISTRIBUTOR VENDOR"`
	CompanyName   string `json:"company_name" binding:"required"`
	ContactPerson string `json:"contact_person"`
	Email         string `json:"email" binding:"omitempty,email"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	Address       string `json:"address"`
	Status        string `json:"status" binding:"omitempty,oneof=ACTIVE INACTIVE BLACKLIST"`
	LogoURL       string `json:"logo_url"`
}
