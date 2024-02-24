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
		NameCity:       NameCity,
		NameCountry:    client.Sys.Country,
		Temperature:    float64(client.Main.Temp),
		MinTemperature: float64(client.Main.TempMin),
		MaxTemperature: float64(client.Main.TempMax),
		Humidity:       uint(client.Main.Humidity),
		Condition:      client.Weather[0].Main,
		Description:    client.Weather[0].Description,
		Wind:           float64(client.Wind.Speed),
		Rain:           float64(client.Rain.OneH),
		Date:           time.Unix(int64(client.Dt), 0),
	}

	return newCity
}
