package entity

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type Location struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
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
	multipler := math.Pow(10, float64(decimals))
	return math.Round(value*multipler) / multipler
}

func (location Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]float64{
		"latitude":  roundToDecimals(location.Lat, 4),
		"longitude": roundToDecimals(location.Long, 4),
	})
}
