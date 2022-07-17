package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/kafka"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func UpdateStatus(c *gin.Context) {
	status, validatedInput := utils.ValidateUpdateStatusInput(c)

	if !status {
		return
	}
	go kafka.SendToConsumer(
		validatedInput.TheaterID,
		validatedInput.TimeSlotID,
		validatedInput.SeatID,
		validatedInput.Status,
	)

	// FIXME: REMOVE THIS LINE; DO NOT DO THIS
	// THIS IS JUST FOR TESTING
	kafka.ReadMessage(kafka.TOPIC)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updating status...",
	})
}
