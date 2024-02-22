package handler

import "time"

type CityRequest struct {
	NameCity string
}

type CityResponse struct {
	NameCity    string
	NameCountry string
	Temperature uint
	Humidity    uint
	Description string
	Date        time.Time
}
