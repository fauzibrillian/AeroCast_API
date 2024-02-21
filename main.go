package main

import (
	"AeroCast_API/config"
	"AeroCast_API/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()
	mongoClient, err := database.InitMongoDB(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}
}
