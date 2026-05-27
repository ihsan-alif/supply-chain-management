package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreateWarehouseHandler(c *gin.Context) {

	var request dto.WarehouseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.CreateWarehouse(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Warehouse created",
	})
}

func GetWarehousesHandler(c *gin.Context) {

	warehouses, err := services.GetWarehouses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": warehouses,
	})
}

func UpdateWarehouseHandler(c *gin.Context) {

	id := c.Param("id")

	var request dto.WarehouseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.UpdateWarehouse(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse updated",
	})
}

func DeleteWarehouseHandler(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteWarehouse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messsage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse deleted",
	})
}
