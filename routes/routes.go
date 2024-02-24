package routes

import (
	"AeroCast_API/features/prediction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, ch prediction.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeCity(e, ch)

}
func routeCity(e *echo.Echo, ch prediction.Handler) {
	e.POST("/city", ch.AddCity())
	e.GET("/city", ch.SearchCity())

}
