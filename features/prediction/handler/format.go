package handler

import "time"

type CityRequest struct {
	NameCity string
}

type CityResponse struct {
	NameCity       string
	NameCountry    string
	Temperature    float64
	MinTemperature float64
	MaxTemperature float64
	Humidity       uint
	Condition      string
	Description    string
	Wind           float64
	Rain           float64
	Date           time.Time
}
