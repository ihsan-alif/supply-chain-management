package repositories

import (
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
)

func CreateUser(user *models.User) error {

	return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := config.DB.
		Preload("Role").
		Where("email = ?", email).
		First(&user).Error

	return &user, err
}