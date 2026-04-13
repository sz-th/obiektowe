package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WeatherRequest struct {
	City string `json:"city" form:"city" query:"city"`
}

type WeatherController struct{}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	city := c.QueryParam("city")
	
	if city == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Podaj miasto w parametrze, np. /weather?city=Krakow",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"location": city,
		"temp":     15.5,
		"status":   "success",
		"method":   "GET",
	})
}

func (wc *WeatherController) PostWeather(c echo.Context) error {
	req := new(WeatherRequest)
	
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Nieprawidłowy format danych",
		})
	}

	if req.City == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Pole 'city' jest wymagane w ciele zapytania JSON",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"location": req.City,
		"temp":     14.0,
		"status":   "success",
		"method":   "POST",
	})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	weatherController := &WeatherController{}

	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.PostWeather)

	e.Logger.Fatal(e.Start(":8080"))
}