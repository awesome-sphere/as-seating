package main

import (
	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server_router := gin.Default()
	server_router.POST("/seating/get-all-seats", service.GetAllSeats)
	server_router.POST("/seating/get-booked-seats", service.GetBookedSeat)
	server_router.POST("/seating/check-seat", service.CheckSeat)
	server_router.POST("/seating/init-seats", service.InitSeats)

	redis.InitializeRedisConn()
	server_router.Run(":9004") // listen and serve on 0.0.0.0:9000
}
