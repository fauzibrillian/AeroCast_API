package prediction

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Prediction struct {
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

type Handler interface {
	AddCity() echo.HandlerFunc
	SearchCity() echo.HandlerFunc
}

type Repository interface {
	NewCity(NewCity Prediction) (Prediction, error)
	SearchCity(NameCity string, NameCountry string, page uint, limit uint) ([]Prediction, uint, error)
}

type Service interface {
	NewCity(NewCity Prediction) (Prediction, error)
	SearchCity(NameCity string, NameCountry string, page uint, limit uint) ([]Prediction, uint, error)
}
