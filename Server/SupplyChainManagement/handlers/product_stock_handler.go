package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreateProductStockHandler(c *gin.Context) {

	var request dto.ProductStockRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	err := services.CreateProductStock(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product stock added",
	})
}

func GetProductStocksHandler(c *gin.Context) {

	productStocks, err := services.GetProductStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productStocks,
	})
}