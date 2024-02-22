package service

import (
	"AeroCast_API/features/prediction"
	"errors"
)

type CityServices struct {
	repo prediction.Repository
}

func New(c prediction.Repository) prediction.Service {
	return &CityServices{
		repo: c,
	}
}

// NewCity implements prediction.Service.
func (cs *CityServices) NewCity(NewCity prediction.Prediction) (prediction.Prediction, error) {
	result, err := cs.repo.NewCity(NewCity)
	if err != nil {
		return prediction.Prediction{}, errors.New("inputan tidak boleh kosong")
	}

	return result, err
}
