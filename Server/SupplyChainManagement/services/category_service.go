package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
)

func CreateCategory(request dto.CategoryRequest) error {

	category := models.Category{
		Name: request.Name,
	}

	return repositories.CreateCategory(&category)
}

func GetCategories(limit, offset int, search string) ([]models.Category, int64, error) {

	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return repositories.GetCategories(limit, offset, search)
}

func GetCategoryByID(id string) (*models.Category, error) {

	categoryID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return repositories.GetCategoryByID(categoryID)
}

func UpdateCategory(id string, request dto.CategoryRequest) error {

	categoryID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = repositories.GetCategoryByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return repositories.UpdateCategoryName(categoryID, request.Name)
}

func DeleteCategory(id string) error {

	categoryID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = repositories.GetCategoryByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return repositories.DeleteCategory(categoryID)
}
