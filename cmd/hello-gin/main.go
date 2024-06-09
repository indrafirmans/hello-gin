package main

import (
	"github.com/indrafirmans/hello-gin/config"
	"github.com/indrafirmans/hello-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.LoadConfig()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
