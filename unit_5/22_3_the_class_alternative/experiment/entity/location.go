package entity

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type Location struct {
	Name        string  `json:"name"`
	LandingSite string  `json:"landingSite"`
	Lat         float64 `json:"latitude"`
	Long        float64 `json:"longitude"`
}

func NewLocation(name, landingSite string, latCoord, longCoord Coordinate) Location {
	return Location{
		Name:        name,
		LandingSite: landingSite,
		Lat:         roundToDecimals(latCoord.Decimal(), 4),
		Long:        roundToDecimals(longCoord.Decimal(), 4),
	}
}

func (location Location) Print() {
	bytes, err := json.MarshalIndent(location, "", "\t")
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func roundToDecimals(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// func (location Location) MarshalJSON() ([]byte, error) {
// 	type Alias Location

// 	return json.Marshal(&struct {
// 		Name        string  `json:"name"`
// 		LandingSite string  `json:"landingSite"`
// 		Lat         float64 `json:"latitude"`
// 		Long        float64 `json:"longitude"`
// 	}{
// 		Name:        location.Name,
// 		LandingSite: location.LandingSite,
// 		Lat:         roundToDecimals(location.Lat, 4),
// 		Long:        roundToDecimals(location.Long, 4),
// 	})
// }
