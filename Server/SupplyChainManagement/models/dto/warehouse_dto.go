package dto

type WarehouseRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
}
