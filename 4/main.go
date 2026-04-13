package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Weather struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	City        string  `gorm:"uniqueIndex" json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

type WeatherRequest struct {
	City string `json:"city" form:"city" query:"city"`
}

type WeatherController struct {
	DB *gorm.DB
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		panic("Nie udało się połączyć z bazą danych")
	}

	db.AutoMigrate(&Weather{})

	var count int64
	db.Model(&Weather{}).Count(&count)

	if count == 0 {
		initialData := []Weather{
			{City: "Warszawa", Temperature: 15.5, Description: "Częściowo zachmurzone"},
			{City: "Krakow", Temperature: 14.0, Description: "Lekki deszcz"},
			{City: "Wroclaw", Temperature: 17.2, Description: "Słonecznie"},
			{City: "Gdansk", Temperature: 12.0, Description: "Wietrznie"},
		}
		
		db.Create(&initialData)
	}

	return db
}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Podaj miasto w parametrze, np. /weather?city=Krakow",
		})
	}

	var weather Weather
	result := wc.DB.Where("LOWER(city) = ?", strings.ToLower(city)).First(&weather)
	
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Brak danych dla miasta: " + city,
		})
	}

	return c.JSON(http.StatusOK, weather)
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
			"error": "Pole 'city' jest wymagane",
		})
	}

	var weather Weather
	result := wc.DB.Where("LOWER(city) = ?", strings.ToLower(req.City)).First(&weather)
	
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Brak danych dla miasta: " + req.City,
		})
	}

	return c.JSON(http.StatusOK, weather)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := initDB()

	weatherController := &WeatherController{
		DB: db,
	}

	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.PostWeather)

	e.Logger.Fatal(e.Start(":8080"))
}