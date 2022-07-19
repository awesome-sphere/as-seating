package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/redis"
	"github.com/gin-gonic/gin"
)

func InitSeats(c *gin.Context) {
	for theater := 1; theater <= 5; theater++ {
		for timeSlot := 1; timeSlot <= 3; timeSlot++ {
			for seat := 1; seat <= 55; seat++ {
				redis.UpdateStatus(int64(theater), int64(timeSlot), seat, redis.AVAILABLE, false)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully initialized seats.",
	})
}
