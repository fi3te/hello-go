package weather

type dataPoint struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature"`
	Rain        float64 `json:"rain"`
	Humidity    float64 `json:"humidity"`
}
