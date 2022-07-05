package main

import (
	"github.com/awesome-sphere/as-seating/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/authentication", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	models.InitDatabase()
	server.Run(":9000") // listen and serve on 0.0.0.0:8080
}
