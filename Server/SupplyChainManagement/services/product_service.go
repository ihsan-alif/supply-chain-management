package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
)

func CreateProduct(request dto.ProductRequest) error {

	categoryUUID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return err
	}

	_, err = repositories.GetCategoryByID(categoryUUID)
	if err != nil {
		return errors.New("category not found")
	}

	product := models.Product{
		CategoryID:   categoryUUID,
		Code:         request.Code,
		Name:         request.Name,
		Description:  request.Description,
		Unit:         request.Unit,
		ImageURL:     request.ImageURL,
		MinimumStock: request.MinimumStock,
	}

	return repositories.CreateProduct(&product)
}

func GetProducts(limit, offset int, search string) ([]models.Product, int64, error) {

	return repositories.GetProducts(
		limit,
		offset,
		search,
	)
}

func GetProductByID(id string) (*models.Product, error) {

	productID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return repositories.GetProductByID(productID)
}

func UpdateProduct(id string, request dto.ProductRequest) error {

	productID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	categoryUUID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return err
	}

	_, err = repositories.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	_, err = repositories.GetCategoryByID(categoryUUID)
	if err != nil {
		return errors.New("category not found")
	}

	product := models.Product{
		CategoryID:   categoryUUID,
		Code:         request.Code,
		Name:         request.Name,
		Description:  request.Description,
		Unit:         request.Unit,
		ImageURL:     request.ImageURL,
		MinimumStock: request.MinimumStock,
	}

	product.ID = productID

	return repositories.UpdateProduct(&product)
}

func DeleteProduct(id string) error {

	productID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = repositories.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	return repositories.DeleteProduct(productID)
}
