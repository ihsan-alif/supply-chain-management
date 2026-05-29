package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreatePartnerHandler(c *gin.Context) {

	var request dto.PartnerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.CreatePartner(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Partner created",
	})
}

func GetPartnersHandler(c *gin.Context) {

	partnerType := c.Query("type")

	partners, err := services.GetPartners(partnerType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": partners,
	})
}

func GetPartnerByIDHandler(c *gin.Context) {

	id := c.Param("id")

	partner, err := services.GetPartnerByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "invalid uuid") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "partner not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": partner,
	})
}

func UpdatePartnerHandler(c *gin.Context) {

	id := c.Param("id")

	var request dto.PartnerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.UpdatePartner(id, request)
	if err != nil {
		errStr := err.Error()

		if strings.Contains(errStr, "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": errStr,
			})
			return
		}

		if strings.Contains(errStr, "invalid uuid") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid uuid format",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "partner updated",
	})
}

func DeletePartnerHandler(c *gin.Context) {

	id := c.Param("id")

	err := services.DeletePartner(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		if strings.Contains(err.Error(), "invalid uuid") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid uuid format",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "partner deleted",
	})
}