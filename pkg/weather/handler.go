package weather

import (
	"encoding/json"
	"net/http"
)

const latitude = 51.515
const longitude = 7.465

func RegisterHttpHandler() {
	http.HandleFunc("/weather", handle)
}

func handle(w http.ResponseWriter, req *http.Request) {
	data, err := fetchForecast(latitude, longitude)
	encoder := json.NewEncoder(w)
	if err == nil {
		encoder.Encode(data)
	} else {
		encoder.Encode(err.Error())
	}
}
