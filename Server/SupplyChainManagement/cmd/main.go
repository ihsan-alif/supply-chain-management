package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/routes"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.AuthRoute(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
