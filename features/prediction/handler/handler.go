package handler

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/helper/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type CityHandler struct {
	h prediction.Service
}

func New(h prediction.Service) prediction.Handler {
	return &CityHandler{
		h: h,
	}
}

// AddCity implements prediction.Handler.
func (ch *CityHandler) AddCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CityRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(prediction.Prediction)
		inputProcess.NameCity = input.NameCity

		result, err := ch.h.NewCity(*inputProcess)
		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			if strings.Contains(err.Error(), "tidak memiliki izin") {
				statusCode = http.StatusNotFound
				message = "product tidak ditemukan"
			} else if strings.Contains(err.Error(), "tidak memiliki izin") {
				statusCode = http.StatusForbidden
				message = "Anda tidak memiliki izin untuk menghapus user ini"
			}

			return responses.PrintResponse(c, statusCode, message, nil)
		}

		var response = new(CityResponse)
		response.NameCity = result.NameCity
		response.NameCountry = result.NameCountry
		response.Temperature = result.Temperature
		response.Humidity = result.Humidity
		response.Description = result.Description
		response.Date = result.Date

		return responses.PrintResponse(c, http.StatusCreated, "Success Create City Data", response)
	}
}
