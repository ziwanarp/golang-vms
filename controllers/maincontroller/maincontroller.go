package maincontroller

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/type/latlng"
)

type json struct {
	Fields struct {
		GeoPoint struct {
			latlng.LatLng
		}
		St struct {
			BooleanValue bool `json:"booleanValue"`
		}
		Vel struct {
			DoubleValue float32 `json:"doubleValue"`
		}
		Ts struct {
			StringValue string `json:"stringValue"`
		}
	}
}

func Post(c *gin.Context) {

	var data json

	//bind data body ke var struct
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

	lokasi := data.Fields.GeoPoint.LatLng

	// update doc returning update doc
	if updateDoc(ctx, client, &data, &lokasi) {
		c.JSON(http.StatusOK, gin.H{"status": "OK",
			"code": 200})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": "BAD",
			"code": 400})
	}
}

func updateDoc(ctx context.Context, client *firestore.Client, data *json, lokasi *latlng.LatLng) bool {

	_, err := client.Collection("DataPerusahaan").Doc("ntsClone").Collection("DaftarKendaraan").Doc("IDKendaraan_001").Update(ctx, []firestore.Update{
		{
			Path:  "GeoPoint",
			Value: lokasi,
		},
		{
			Path:  "St",
			Value: data.Fields.St.BooleanValue,
		},
		{
			Path:  "Vel",
			Value: data.Fields.Vel.DoubleValue,
		},
		{
			Path:  "Ts",
			Value: data.Fields.Ts.StringValue,
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return true
}
