package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func RegisterHandler(c *gin.Context) {

	var request dto.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register success",
	})
}

func LoginHandler(c *gin.Context) {

	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := services.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}