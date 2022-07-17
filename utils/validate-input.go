package utils

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/gin-gonic/gin"
)

func ValidateInput(c *gin.Context) (bool, serializer.SeatsInputSerializer) {
	var input_serializer serializer.SeatsInputSerializer
	if err := c.BindJSON(&input_serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.SeatsInputSerializer{}
	}
	return true, input_serializer
}

func ValidateCheckSeatInput(c *gin.Context) (bool, serializer.CheckSeatInputSerializer) {
	var input_serializer serializer.CheckSeatInputSerializer
	if err := c.BindJSON(&input_serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.CheckSeatInputSerializer{}
	}
	return true, input_serializer
}

func ValidateUpdateStatusInput(c *gin.Context) (bool, serializer.SeatModelSerializer) {
	var input_serializer serializer.SeatModelSerializer
	if err := c.BindJSON(&input_serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.SeatModelSerializer{}
	}
	return true, input_serializer
}
