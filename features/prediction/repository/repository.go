package repository

import (
	"AeroCast_API/features/prediction"
	"AeroCast_API/helper/openweather"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		Date:           weather.Date,
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

// SearchCity implements prediction.Repository.
func (cq *CityQuery) SearchCity(NameCity string, NameCountry string, page uint, limit uint) ([]prediction.Prediction, uint, error) {
	offset := (page - 1) * limit
	filter := bson.M{}
	if NameCity != "" {
		filter["name_city"] = NameCity
	}
	if NameCountry != "" {
		filter["name_country"] = NameCountry
	}

	options := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.M{"date": -1})

	cursor, err := cq.db.Collection(cq.collection).Find(context.Background(), filter, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var predictions []prediction.Prediction
	for cursor.Next(context.Background()) {
		var cityModel CityModel
		if err := cursor.Decode(&cityModel); err != nil {
			return nil, 0, err
		}

		predictions = append(predictions, prediction.Prediction{
			NameCity:       cityModel.NameCity,
			NameCountry:    cityModel.NameCountry,
			Temperature:    cityModel.Temperature,
			MinTemperature: cityModel.MinTemperature,
			MaxTemperature: cityModel.MaxTemperature,
			Humidity:       cityModel.Humidity,
			Condition:      cityModel.Condition,
			Description:    cityModel.Description,
			Wind:           cityModel.Wind,
			Rain:           cityModel.Rain,
			Date:           cityModel.Date,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	totalCount, err := cq.db.Collection(cq.collection).CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}

	return predictions, uint(totalCount), nil
}
