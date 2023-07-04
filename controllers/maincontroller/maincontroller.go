package maincontroller

import (
	// "encoding/json"
	"log"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type data struct {
	field string
	Lat   string
	Lon   string
	St    string
	Vel   string
	Ts    string
}

func Post(c *gin.Context) {
	var newData data

	log.Println(newData)
	if err := c.ShouldBindJSON(newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println(newData)
	c.JSON(http.StatusOK, gin.H{"fields": newData})
}
