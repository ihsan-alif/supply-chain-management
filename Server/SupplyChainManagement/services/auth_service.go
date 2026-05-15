package services

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
	"github.com/ihsan-alif/supply-chain-management/utils"
)

func Register(request dto.RegisterRequest) error {

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	roleUUID, err := uuid.Parse(request.RoleID)
	if err != nil {
		return err
	}

	user := models.User{
		RoleID:   roleUUID,
		FullName: request.FullName,
		Email:    request.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	return repositories.CreateUser(&user)
}

func Login(request dto.LoginRequest) (string, error) {

	user, err := repositories.FindUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	validPassword := utils.CheckPassword(
		request.Password,
		user.Password,
	)
	if !validPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}