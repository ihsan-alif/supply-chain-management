package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/handlers"
	"github.com/ihsan-alif/supply-chain-management/middleware"
)

func AuthRoute(r *gin.Engine) {

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/login", handlers.LoginHandler)

	}

	authProtected := r.Group("/api")
	authProtected.Use(middleware.AuthMiddleware())
	{
		authProtected.GET("/profile", handlers.ProfileHandler)
		authProtected.GET("/me", handlers.Me)
	}

	admin := authProtected.Group("/admin")
	admin.Use(middleware.RoleMiddleware("ADMIN"))
	admin.GET("/dashboard", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Welcome Admin",
		})
	})
}
