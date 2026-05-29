package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func StockMovementRoute(r *gin.Engine) {

	movement := r.Group("/api/stock-movements")

	movement.Use(middleware.AuthMiddleware())

	movement.POST("/", handlers.CreateStockMovementHandler)
	movement.GET("/", handlers.GetStockMovementsHandler)
}
