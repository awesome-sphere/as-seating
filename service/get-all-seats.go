package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func GetAllSeats(c *gin.Context) {
	status, validatedInput := utils.ValidateInput(c)
	if !status {
		return
	}

	output, err := redis.ScanPrefix(validatedInput.TheaterID, validatedInput.TimeSlotID, "")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error while getting seats.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "Successfully checked seats.",
		"seats_info": output,
	})
}
