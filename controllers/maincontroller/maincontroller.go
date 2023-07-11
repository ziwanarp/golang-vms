package maincontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
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

		// Use a service account
		ctx := context.Background()
		sa := option.WithCredentialsFile("auth.json")
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			log.Fatalln(err)
		}
	
		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()
	

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"status": "OK",
		"code": 200,
		"data": data})
}
