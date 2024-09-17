package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const url = "https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&hourly=temperature_2m,relativehumidity_2m,rain"

type openMeteoData struct {
	Elevation    int64
	Hourly_units struct {
		Relativehumidity_2m string
		Rain                string
		Temperature_2m      string
		Time                string
	}
	Latitude           float64
	Utc_offset_seconds int64
	Generationtime_ms  float64
	Longitude          float64
	Hourly             struct {
		Temperature_2m      []float64
		Rain                []float64
		Time                []string // e.g. 2022-06-26T00:00
		Relativehumidity_2m []float64
	}
}

func fetchForecast(latitude float64, longitude float64) ([]dataPoint, error) {
	resp, err := http.Get(fmt.Sprintf(url, latitude, longitude))
	var openMeteoData openMeteoData
	var dataPoints []dataPoint
	if err == nil {
		err = json.NewDecoder(resp.Body).Decode(&openMeteoData)
		dataPoints, err = toDataPoints(openMeteoData)
	}
	return dataPoints, err
}

func toDataPoints(data openMeteoData) ([]dataPoint, error) {
	series := data.Hourly
	length1 := len(series.Temperature_2m)
	length2 := len(series.Rain)
	length3 := len(series.Time)
	length4 := len(series.Relativehumidity_2m)
	length := length1
	if length1 != length2 || length2 != length3 || length3 != length4 {
		return []dataPoint{}, errors.New("Length of series are not equal.")
	} else {
		result := []dataPoint{}
		for i := 0; i < length; i++ {
			result = append(result, dataPoint{
				series.Time[i],
				series.Temperature_2m[i],
				series.Rain[i],
				series.Relativehumidity_2m[i],
			})
		}
		return result, nil
	}
}
