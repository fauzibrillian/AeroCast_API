package handler

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/helper/responses"
	"net/http"
	"strconv"
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
		response.MinTemperature = result.MinTemperature
		response.MaxTemperature = result.MaxTemperature
		response.Humidity = result.Humidity
		response.Condition = result.Condition
		response.Description = result.Description
		response.Wind = result.Wind
		response.Rain = result.Rain
		response.Date = result.Date

		return responses.PrintResponse(c, http.StatusCreated, "Success Create City Data", response)
	}
}

// SearchCity implements prediction.Handler.
func (ch *CityHandler) SearchCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		NameCity := c.QueryParam("NameCity")
		NameCountry := c.QueryParam("NameCountry")
		uintPage := uint(page)
		uintLimit := uint(limit)

		city, totalPage, err := ch.h.SearchCity(NameCity, NameCountry, uintPage, uintLimit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		var response []CityResponse
		for _, result := range city {
			response = append(response, CityResponse{
				Date:           result.Date,
				NameCity:       result.NameCity,
				NameCountry:    result.NameCountry,
				Temperature:    result.Temperature,
				MinTemperature: result.MinTemperature,
				MaxTemperature: result.MaxTemperature,
				Humidity:       result.Humidity,
				Condition:      result.Condition,
				Description:    result.Description,
				Wind:           result.Wind,
				Rain:           result.Rain,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "Get City Successful",
			"data":       response,
			"pagination": map[string]interface{}{"page": page, "limit": limit, "total_page": totalPage},
		})
	}
}
