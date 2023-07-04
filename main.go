package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ziwanarp/vms-go/controllers/maincontroller"
)

func main() {

	router := gin.Default()

	router.GET("/api/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "200 OK",
		})
	})

	router.POST("/api/post", maincontroller.Post)

	router.Run()

}
