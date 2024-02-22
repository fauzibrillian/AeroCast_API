package prediction

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Prediction struct {
	NameCity    string
	NameCountry string
	Temperature uint
	Humidity    uint
	Description string
	Date        time.Time
}

type Handler interface {
	AddCity() echo.HandlerFunc
}

type Repository interface {
	NewCity(NewCity Prediction) (Prediction, error)
}

type Service interface {
	NewCity(NewCity Prediction) (Prediction, error)
}
