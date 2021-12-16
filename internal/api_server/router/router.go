package router

import (
	"github.com/BulatKarimov/gin_rest_api/internal/api_server/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/events", controllers.FindEvents)
	r.POST("/events", controllers.CreateEvent)
}
