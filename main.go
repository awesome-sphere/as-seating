package main

import (
	"github.com/awesome-sphere/as-seating/kafka"
	"github.com/awesome-sphere/as-seating/redis"
	"github.com/awesome-sphere/as-seating/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/seating/get-all-seats", service.GetAllSeats)
	router.GET("/seating/get-booked-seats", service.GetBookedSeat)
	router.GET("/seating/check-seat", service.CheckSeat)

	router.POST("/seating/update-status", service.UpdateStatus)

	redis.InitializeRedisConn()
	kafka.InitKafkaTopic()

	router.Run(":9004") // listen and serve on 0.0.0.0:9000
}
