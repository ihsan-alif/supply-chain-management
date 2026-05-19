package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func ProductRoute(r *gin.Engine) {

	product := r.Group("/api/products")

	product.Use(middleware.AuthMiddleware())

	product.POST("/", handlers.CreateProductHandler)
	product.GET("/", handlers.GetProductsHandler)
	product.GET("/:id", handlers.GetProductByIDHandler)
	product.PUT("/:id", handlers.UpdateProductHandler)
	product.DELETE("/:id", handlers.DeleteProductHandler)
}