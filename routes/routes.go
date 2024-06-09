package routes

import (
	"github.com/indrafirmans/hello-gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.CheckHealth)
	router.GET("/users", controllers.GetUsers)
}
