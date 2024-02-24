package service_test

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/features/prediction/mocks"
	"AeroCast_API/features/prediction/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCity(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	t.Run("Success Case", func(t *testing.T) {
		expectedBooks := []prediction.Prediction{{NameCity: "Sidoarjo", NameCountry: "ID"}, {NameCity: "Nagasaki", NameCountry: "JP"}}
		expectedTotalPage := uint(3)

		repo.On("SearchCity", "Sidoarjo", "ID", uint(1), uint(10)).Return(expectedBooks, expectedTotalPage, nil).Once()
		city, totalPage, err := m.SearchCity("Sidoarjo", "ID", uint(1), uint(10))
		assert.NoError(t, err)
		assert.Equal(t, expectedBooks, city)
		assert.Equal(t, expectedTotalPage, totalPage)
	})

	t.Run("Failed Case", func(t *testing.T) {
		repo.On("SearchCity", "Sidoarjo", "ID", uint(1), uint(10)).Return([]prediction.Prediction{}, uint(0), errors.New("repository error")).Once()
		city, totalPage, err := m.SearchCity("Sidoarjo", "ID", uint(1), uint(10))
		assert.Error(t, err)
		assert.Equal(t, []prediction.Prediction{}, city)
		assert.Equal(t, uint(0), totalPage)
		assert.Equal(t, "repository error", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestNewCity(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	t.Run("Success Case", func(t *testing.T) {
		input := prediction.Prediction{
			NameCity:       "sidoarjo",
			NameCountry:    "ID",
			Temperature:    28.9,
			MinTemperature: 25.6,
			MaxTemperature: 28.9,
			Humidity:       30,
			Condition:      "Haze",
			Description:    "few clouds",
			Wind:           2.56,
			Rain:           0,
		}
		repo.On("NewCity", input).Return(input, nil).Once()
		city, err := m.NewCity(input)

		assert.NoError(t, err, city)
		assert.Equal(t, prediction.Prediction{
			NameCity:       "sidoarjo",
			NameCountry:    "ID",
			Temperature:    28.9,
			MinTemperature: 25.6,
			MaxTemperature: 28.9,
			Humidity:       30,
			Condition:      "Haze",
			Description:    "few clouds",
			Wind:           2.56,
			Rain:           0,
		}, city)

		repo.AssertExpectations(t)
	})
}
