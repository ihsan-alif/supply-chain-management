package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func WarehouseRoute(r *gin.Engine) {

	warehouse := r.Group("/api/warehouses")

	warehouse.Use(middleware.AuthMiddleware())

	warehouse.POST("/", handlers.CreateWarehouseHandler)
	warehouse.GET("/", handlers.GetWarehousesHandler)
	warehouse.PUT("/:id", handlers.UpdateWarehouseHandler)
	warehouse.DELETE("/:id", handlers.DeleteWarehouseHandler)
}
