package service

import (
	"fmt"
	"net/http"

	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func GetBookedSeat(c *gin.Context) {
	status, validated_input := utils.ValidateInput(c)
	if status {
		fmt.Println(models.Available)
		var querySet []models.SeatInfo
		var selectedColumnQuerySet []serializer.SeatOutputSerializer
		err := models.DB.Model(&models.SeatInfo{}).Where(map[string]interface{}{
			"theater_id": validated_input.TheaterID, "status": models.Booked,
		}).Find(&querySet).Find(&selectedColumnQuerySet).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"seats_info": selectedColumnQuerySet,
		})
	}
	return
}
