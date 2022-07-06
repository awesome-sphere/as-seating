package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func GetAllSeats(c *gin.Context) {
	status, validated_input := utils.ValidateInput(c)
	if status {
		var querySet []models.SeatInfo
		var selectedColumnQuerySet []serializer.SeatOutputSerializer
		err := models.DB.Model(&models.SeatInfo{}).Where(
			"theater_id", validated_input.TheaterID,
		).Order("id asc").Find(&querySet).Find(&selectedColumnQuerySet).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		// fmt.Printf("%#v\n", test)
		c.JSON(http.StatusOK, gin.H{
			"seats_info": selectedColumnQuerySet,
		})
	}
	return
}
