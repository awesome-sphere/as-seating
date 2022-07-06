package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/awesome-sphere/as-seating/utils"
	"github.com/gin-gonic/gin"
)

func CheckSeat(c *gin.Context) {
	status, validated_input := utils.ValidateCheckSeatInput(c)
	if status {
		var querySet models.SeatInfo
		var output serializer.CheckSeatOutputSerializer
		err := models.DB.Model(&models.SeatInfo{}).Where(map[string]interface{}{
			"theater_id": validated_input.TheaterID, "id": validated_input.SeatID,
		}).First(&querySet).Find(&output).Error
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{
					"status": "NOT FOUND",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		c.JSON(http.StatusOK, output)
	}
	return
}
