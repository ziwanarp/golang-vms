package mainmodel

import "google.golang.org/genproto/googleapis/type/latlng"

type Receipt struct {
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
