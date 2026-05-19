package dto

type ProductRequest struct {
	CategoryID   string `json:"category_id" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Unit         string `json:"unit" binding:"required"`
	ImageURL     string `json:"image_url"`
	MinimumStock int    `json:"minimum_stock"`
}
