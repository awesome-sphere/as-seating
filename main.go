package main

import (
	"github.com/awesome-sphere/as-seating/models"
	"github.com/awesome-sphere/as-seating/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server_router := gin.Default()
	server_router.GET("/seating/get_all_seats", service.GetAllSeats)
	server_router.GET("/seating/get_booked_seats", service.GetBookedSeat)
	server_router.GET("/seating/check_seat", service.CheckSeat)
	models.InitDatabase()
	server_router.Run(":9000") // listen and serve on 0.0.0.0:9000
}
