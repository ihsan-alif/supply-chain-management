package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/services"
)

func CreateCategoryHandler(c *gin.Context) {

	var request dto.CategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.CreateCategory(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created",
	})
}

func GetCategoriesHandler(c *gin.Context) {

	page, _ := strconv.Atoi(
		c.DefaultQuery("page", "1"),
	)

	limit, _ := strconv.Atoi(
		c.DefaultQuery("limit", "10"),
	)

	search := c.Query("search")

	offset := (page - 1) * limit

	categories, total, err := services.GetCategories(
		limit,
		offset,
		search,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetCategoryByIDHandler(c *gin.Context) {

	id := c.Param("id")

	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func UpdateCategoryHandler(c *gin.Context) {

	id := c.Param("id")

	var request dto.CategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := services.UpdateCategory(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated",
	})
}

func DeleteCategoryHandler(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted",
	})
}
