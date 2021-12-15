package controllers

import (
	"net/http"

	"github.com/BulatKarimov/gin_rest_api/app/api_server/models"
	"github.com/gin-gonic/gin"
)

type CreateEventInput struct {
	Type string `json:"type" binding:"required"`
}

func FindEvents(c *gin.Context) {
	var events []models.Event

	models.DB.Find(&events)

	c.JSON(http.StatusOK, gin.H{"data": events})
}

func CreateEvent(c *gin.Context) {
	// Validate input
	var input CreateEventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	event := models.Event{Type: input.Type}
	models.DB.Create(&event)

	c.JSON(http.StatusOK, gin.H{"data": event})
}
