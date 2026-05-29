package repositories

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
)

func CreatePartner(partner *models.Partner) error {

	return config.DB.Create(partner).Error
}

func GetPartners(partnerType string) ([]models.Partner, error) {

	var partners []models.Partner

	query := config.DB.Model(&models.Partner{})

	if partnerType != "" {

		query = query.Where(
			"type = ?",
			partnerType,
		)
	}

	err := query.
		Order("created_at DESC").
		Find(&partners).Error

	return partners, err
}

func GetPartnerByID(id uuid.UUID) (*models.Partner, error) {

	var partner models.Partner

	err := config.DB.
		Where("id = ?", id).
		First(&partner).Error

	return &partner, err
}

func UpdatePartner(partner *models.Partner) error {

	return config.DB.Model(partner).Updates(map[string]interface{}{
		"code":           partner.Code,
		"type":           partner.Type,
		"company_name":   partner.CompanyName,
		"contact_person": partner.ContactPerson,
		"email":          partner.Email,
		"phone":          partner.Phone,
		"city":           partner.City,
		"address":        partner.Address,
		"status":         partner.Status,
		"logo_url":       partner.LogoURL,
	}).Error
}

func DeletePartner(id uuid.UUID) error {

	return config.DB.Delete(
		&models.Partner{},
		"id = ?",
		id,
	).Error
}
