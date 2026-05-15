package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
)

func AuthRoute(r *gin.Engine) {

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/login", handlers.LoginHandler)
	}

}
