package openweather

import (
	"AeroCast_API/config"
	"AeroCast_API/features/prediction"
	"log"

	owm "github.com/briandowns/openweathermap"
)

func OpenWeatherMap(NameCity string) (newCity prediction.Prediction) {
	client, err := owm.NewCurrent("C", "EN", config.InitConfig().ACCUWEATHER_KEY)
	if err != nil {
		log.Println("Error creating OpenWeatherMap client:", err)
		return newCity
	}
	err = client.CurrentByName(NameCity)
	if err != nil {
		log.Println("Error fetching weather data:", err)
		return newCity
	}

	return newCity
}
