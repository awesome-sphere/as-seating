package service

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/gin-gonic/gin"
)

func validateInput(c *gin.Context) (bool, serializer.GetAllSeatsSerializer) {
	var input_serializer serializer.GetAllSeatsSerializer
	if err := c.BindJSON(&input_serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.GetAllSeatsSerializer{}
	}
	return true, input_serializer
}

func GetAllSeats(c *gin.Context) {
	status, validated_input := validateInput(c)
	if status {
		var test []models.SeatInfo
		err := models.DB.Model(&models.SeatInfo{}).Where("theater_id", validated_input.TheaterID).Find(&test).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		// fmt.Printf("%#v\n", test)
		c.JSON(http.StatusOK, gin.H{
			"test": test,
		})
	}
	return
}
