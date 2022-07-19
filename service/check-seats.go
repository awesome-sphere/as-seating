package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func CheckSeat(c *gin.Context) {
	status, validated_input := utils.ValidateCheckSeatInput(c)

	if status {
		seatStatus, err := redis.ReadStatus(validated_input)

		if err != nil {
			redis.UpdateStatus(
				validated_input.TheaterID,
				validated_input.TimeSlotID,
				int(validated_input.SeatID),
				redis.AVAILABLE,
				false,
			)
			seatStatus = redis.AVAILABLE
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  seatStatus,
			"message": "Successfully checked seat.",
		})
	}
}
