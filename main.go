package main

import (
	"fmt"
	"github.com/be-edu/weather-2/v2/location"
	"github.com/be-edu/weather-2/v2/weather"
)

func main() {
	fmt.Printf("Weather package version: %s\n", weather.GetVersion())
	makePrediction(51.509865, -0.118092, "London")
}

func makePrediction(lat float64, long float64, locationName string) error {
	gpsCoords := location.GPSCoords([2]float64{lat, long})
	pred, err := weather.Predict(gpsCoords)

	if err != nil {
		return err
	}

	fmt.Printf("The weather prediction for %s is: %v\n", locationName, pred.ToString())

	return nil
}
