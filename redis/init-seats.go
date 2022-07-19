package redis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitSeats(c *gin.Context) {
	for theater := 1; theater <= 5; theater++ {
		for timeSlot := 1; timeSlot <= 15; timeSlot++ {
			for seat := 1; seat <= 55; seat++ {
				UpdateStatus(int64(theater), int64(timeSlot), seat, AVAILABLE, false)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully initialized seats.",
	})
}
