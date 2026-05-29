package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreateStockMovementHandler(c *gin.Context) {

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
		errStr := err.Error()
		if errStr == "insufficient stock" || errStr == "invalid stock movement type" || errStr == "insufficient stock: data stock does not exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errStr,
			})
			return
		}

		if strings.Contains(errStr, "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": errStr,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Stock movement created",
	})
}

func GetStockMovementsHandler(c *gin.Context) {

	movements, err := services.GetStockMovements()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movements,
	})
}
