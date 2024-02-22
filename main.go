package main

import (
	"AeroCast_API/config"
	"AeroCast_API/database"
	ch "AeroCast_API/features/prediction/handler"
	cr "AeroCast_API/features/prediction/repository"
	cs "AeroCast_API/features/prediction/service"
	"AeroCast_API/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	mongoClient, err := database.InitMongoDB(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	cityRepo := cr.New(mongoClient.Client().Database(cfg.DBNAME), "cities")
	cityService := cs.New(cityRepo)
	cityHandler := ch.New(cityService)

	routes.InitRoute(e, cityHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
