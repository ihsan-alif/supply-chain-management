package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
)

func CreatePartner(request dto.PartnerRequest) error {
	
	status := request.Status
	if status == "" {
		request.Status = "ACTIVE"
	}

	partner := models.Partner{
		Code:          request.Code,
		Type:          request.Type,
		CompanyName:   request.CompanyName,
		ContactPerson: request.ContactPerson,
		Email:         request.Email,
		Phone:         request.Phone,
		City:          request.City,
		Address:       request.Address,
		Status:        status,
		LogoURL:       request.LogoURL,
	}

	return repositories.CreatePartner(&partner)
}

func GetPartners(partnerType string) ([]models.Partner, error) {

	return repositories.GetPartners(partnerType)
}

func GetPartnerByID(id string) (*models.Partner, error) {

	partnerID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return repositories.GetPartnerByID(partnerID)
}

func UpdatePartner(id string, request dto.PartnerRequest) error {

	partnerID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = repositories.GetPartnerByID(partnerID)
	if err != nil {
		return errors.New("partner not found")
	}

	status := request.Status
	if status == "" {
		request.Status = "ACTIVE"
	}

	partner := models.Partner{
		Code:          request.Code,
		Type:          request.Type,
		CompanyName:   request.CompanyName,
		ContactPerson: request.ContactPerson,
		Email:         request.Email,
		Phone:         request.Phone,
		City:          request.City,
		Address:       request.Address,
		Status:        status,
		LogoURL:       request.LogoURL,
	}

	partner.ID = partnerID

	return repositories.UpdatePartner(&partner)
}

func DeletePartner(id string) error {

	partnerID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = repositories.GetPartnerByID(partnerID)
	if err != nil {
		return errors.New("partner not found")
	}

	return repositories.DeletePartner(partnerID)
}