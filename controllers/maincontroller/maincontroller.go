package maincontroller

import (
	// "encoding/json"
	// "log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type json struct {
	Fields struct {
		Lat struct {
			DoubleValue string `json:"doubleValue"`
		}
		Lon struct {
			DoubleValue string `json:"doubleValue"`
		}
		St struct {
			BooleanValue string `json:"booleanValue"`
		}
		Vel struct {
			DoubleValue string `json:"doubleValue"`
		}
		Ts struct {
			StringValue string `json:"stringValue"`
		}
	}
}

func Post(c *gin.Context) {

	var data json

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"status": "OK",
		"code": 200,
		"data": data})
}
