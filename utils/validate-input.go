package utils

import (
	"net/http"

	"github.com/awesome-sphere/as-seating/serializer"
	"github.com/gin-gonic/gin"
)

func ValidateInput(c *gin.Context) (bool, serializer.SeatsInputSerializer) {
	var inputSerializer serializer.SeatsInputSerializer
	if err := c.BindJSON(&inputSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.SeatsInputSerializer{}
	}
	return true, inputSerializer
}

func ValidateCheckSeatInput(c *gin.Context) (bool, serializer.CheckSeatInputSerializer) {
	var inputSerializer serializer.CheckSeatInputSerializer
	if err := c.BindJSON(&inputSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.CheckSeatInputSerializer{}
	}
	return true, inputSerializer
}

func ValidateUpdateStatusInput(c *gin.Context) (bool, serializer.SeatModelSerializer) {
	var inputSerializer serializer.SeatModelSerializer
	if err := c.BindJSON(&inputSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false, serializer.SeatModelSerializer{}
	}
	return true, inputSerializer
}
