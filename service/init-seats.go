package service

import (
	"fmt"
	"net/http"

	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func InitSeats(c *gin.Context) {
	status, validated_input := utils.ValidateInput(c)
	if !status {
		return
	}
	for i := 1; i <= 55; i++ {
		_, err := redis.CLIENT.Set(fmt.Sprintf(
			"%d-%d-%d",
			validated_input.TheaterID,
			validated_input.TimeSlotID,
			i),
			redis.AVAILABLE,
			0).Result()
		if err != nil {
			panic(err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully initialized seats.",
	})
}
