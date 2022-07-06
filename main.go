package main

import (
	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server_router := gin.Default()
	server_router.POST("/seating/get-all-seats", service.GetAllSeats)
	server_router.POST("/seating/get-booked-seats", service.GetBookedSeat)
	server_router.POST("/seating/check-seat", service.CheckSeat)
	models.InitDatabase()
	server_router.Run(":9000") // listen and serve on 0.0.0.0:9000
}
