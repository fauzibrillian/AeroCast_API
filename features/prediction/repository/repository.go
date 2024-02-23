package repository

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/helper/openweather"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityModel struct {
	NameCity       string    `bson:"name_city"`
	NameCountry    string    `bson:"name_country"`
	Temperature    float64   `bson:"temperature"`
	MinTemperature float64   `bson:"min_temperature"`
	MaxTemperature float64   `bson:"max_temperature"`
	Humidity       uint      `bson:"humidity"`
	Condition      string    `bson:"condition"`
	Description    string    `bson:"description"`
	Wind           float64   `bson:"wind"`
	Rain           float64   `bson:"rain"`
	Date           time.Time `bson:"date"`
}

type CityQuery struct {
	db         *mongo.Database
	collection string
}

func New(client *mongo.Database, collection string) prediction.Repository {
	return &CityQuery{
		db:         client,
		collection: collection,
	}
}

// NewCity implements prediction.Repository.
func (cq *CityQuery) NewCity(NewCity prediction.Prediction) (prediction.Prediction, error) {
	weather := openweather.OpenWeatherMap(NewCity.NameCity)

	// Create a new CityModel instance
	inputDB := &CityModel{
		NameCity:       NewCity.NameCity,
		NameCountry:    weather.NameCountry,
		Temperature:    weather.Temperature,
		MinTemperature: weather.MinTemperature,
		MaxTemperature: weather.MaxTemperature,
		Humidity:       weather.Humidity,
		Condition:      weather.Condition,
		Description:    weather.Description,
		Wind:           weather.Wind,
		Rain:           weather.Rain,
		Date:           time.Now(),
	}

	// Insert the CityModel into the MongoDB collection
	_, dbInsertErr := cq.db.Collection(cq.collection).InsertOne(context.Background(), inputDB)
	if dbInsertErr != nil {
		return prediction.Prediction{}, dbInsertErr
	}

	// Return the updated newCity with the fetched weather information
	NewCity.NameCountry = inputDB.NameCountry
	NewCity.Temperature = inputDB.Temperature
	NewCity.MinTemperature = inputDB.MinTemperature
	NewCity.MaxTemperature = inputDB.MaxTemperature
	NewCity.Humidity = inputDB.Humidity
	NewCity.Condition = inputDB.Condition
	NewCity.Description = inputDB.Description
	NewCity.Wind = inputDB.Wind
	NewCity.Rain = inputDB.Rain
	NewCity.Date = inputDB.Date

	return NewCity, nil
}
