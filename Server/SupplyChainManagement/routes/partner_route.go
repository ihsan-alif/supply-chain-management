package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func PartnerRoute(r *gin.Engine) {

	partners := r.Group("/api/partners")

	partners.Use(middleware.AuthMiddleware())
	{
		partners.POST("/", handlers.CreatePartnerHandler)
		partners.GET("/", handlers.GetPartnersHandler)
		partners.GET("/:id", handlers.GetPartnerByIDHandler)
		partners.PUT("/:id", handlers.UpdatePartnerHandler)
		partners.DELETE("/:id", handlers.DeletePartnerHandler)
	}
}
