package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func CategoryRoute(r *gin.Engine) {

	category := r.Group("/api/categories")

	category.Use(middleware.AuthMiddleware())

	category.POST("/", handlers.CreateCategoryHandler)
	category.GET("/", handlers.GetCategoriesHandler)
	category.GET("/:id", handlers.GetCategoryByIDHandler)
	category.PUT("/:id", handlers.UpdateCategoryHandler)
	category.DELETE("/:id", handlers.DeleteCategoryHandler)
}