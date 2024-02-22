package repository

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/helper/openweather"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityModel struct {
	NameCity    string    `bson:"name_city"`
	NameCountry string    `bson:"name_country"`
	Temperature uint      `bson:"temperature"`
	Humidity    uint      `bson:"humidity"`
	Description string    `bson:"description"`
	Date        time.Time `bson:"date"`
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
		NameCity:    NewCity.NameCity,
		NameCountry: weather.NameCountry,
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		Description: weather.Description,
		Date:        time.Now(),
	}

	// Insert the CityModel into the MongoDB collection
	_, dbInsertErr := cq.db.Collection(cq.collection).InsertOne(context.Background(), inputDB)
	if dbInsertErr != nil {
		return prediction.Prediction{}, dbInsertErr
	}

	// Return the updated newCity with the fetched weather information
	NewCity.Temperature = inputDB.Temperature
	NewCity.Humidity = inputDB.Humidity
	NewCity.Description = inputDB.Description
	NewCity.Date = inputDB.Date

	return NewCity, nil
}
