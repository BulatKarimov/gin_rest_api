package main

import (
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/controllers"
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/events", controllers.FindEvents)
	r.POST("/events", controllers.CreateEvent)

	r.Run()
}
