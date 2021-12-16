package main

import (
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/controllers"
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/models"
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/router"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	r.GET("/events", controllers.FindEvents)
	r.POST("/events", controllers.CreateEvent)
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	router.InitializeRoutes(r)

	r.Run()
}
