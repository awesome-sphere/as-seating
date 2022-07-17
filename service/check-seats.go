package service

import (
	"fmt"
	"net/http"

	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func CheckSeat(c *gin.Context) {
	status, validated_input := utils.ValidateCheckSeatInput(c)

	if status {
		seat_status, err := redis.CLIENT.Get(fmt.Sprintf(
			"%d-%d-%d",
			validated_input.TheaterID,
			validated_input.TimeSlotID,
			validated_input.SeatID,
		)).Result()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "This seat does not exist.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  seat_status,
			"message": "Successfully checked seat.",
		})
	}
}
