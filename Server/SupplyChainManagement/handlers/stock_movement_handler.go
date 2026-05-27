package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreateStockMovement(c *gin.Context) {

	var request dto.StockMovementRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetString("user_id")

	err := services.CreateStockMovement(request, userID)
	if err != nil {
		if err.Error() == "insufficient stock" || err.Error() == "invalid stock movement type" || err.Error() == "insufficient stock: data stock does not exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Stock movement created",
	})
}

func GetStockMovements(c *gin.Context) {

	movements, err := services.GetStockMovements()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movements,
	})
}
