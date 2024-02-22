package openweather

import (
	"AeroCast_API/config"
	"AeroCast_API/features/prediction"
	"log"
	"time"

	owm "github.com/briandowns/openweathermap"
)

func OpenWeatherMap(NameCity string) (newCity prediction.Prediction) {
	client, err := owm.NewCurrent("C", "EN", config.InitConfig().ACCUWEATHER_KEY)
	if err != nil {
		log.Println("Error creating OpenWeatherMap client:", err)
		return newCity
	}
	client.CurrentByName(NameCity)
	newCity = prediction.Prediction{
		NameCity:    NameCity,
		NameCountry: client.Sys.Country,
		Temperature: uint(client.Main.Temp),
		Humidity:    uint(client.Main.Humidity),
		Description: client.Weather[0].Description,
		Date:        time.Unix(int64(client.Dt), 0),
	}

	return newCity
}
