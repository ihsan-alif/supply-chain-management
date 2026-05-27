package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func ProductStockRoute(r *gin.Engine) {

	productStock := r.Group("/api/product-stocks")

	productStock.Use(middleware.AuthMiddleware())

	productStock.POST("/", handlers.CreateProductStockHandler)
	productStock.GET("/", handlers.GetProductStocksHandler)
}