package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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


type WttrResponse struct {
	CurrentCondition []struct {
		TempC       string `json:"temp_C"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
	} `json:"current_condition"`
}

type WeatherProxy struct{}

func (wp *WeatherProxy) FetchWeather(city string) (*Weather, error) {
	escapedCity := url.QueryEscape(city)
	apiUrl := fmt.Sprintf("https://wttr.in/%s?format=j1", escapedCity)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("błąd zewnętrznego API, status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var wttrData WttrResponse
	if err := json.Unmarshal(body, &wttrData); err != nil {
		return nil, err
	}

	if len(wttrData.CurrentCondition) == 0 {
		return nil, fmt.Errorf("brak danych pogodowych dla miasta: %s", city)
	}

	current := wttrData.CurrentCondition[0]
	tempFloat, _ := strconv.ParseFloat(current.TempC, 64)
	desc := ""
	if len(current.WeatherDesc) > 0 {
		desc = current.WeatherDesc[0].Value
	}

	return &Weather{
		City:        strings.Title(strings.ToLower(city)), 
		Temperature: tempFloat,
		Description: desc,
	}, nil
}


type WeatherController struct {
	DB    *gorm.DB
	Proxy *WeatherProxy 
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Podaj miasto w parametrze"})
	}

	var weather Weather
	result := wc.DB.Where("LOWER(city) = ?", strings.ToLower(city)).First(&weather)
	
	if result.Error != nil {
		proxyData, err := wc.Proxy.FetchWeather(city)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Nie znaleziono w bazie, a zewnętrzne API zwróciło błąd: " + err.Error(),
			})
		}
		
		return c.JSON(http.StatusOK, map[string]interface{}{
			"source": "Zewnętrzne API (Proxy)",
			"data":   proxyData,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"source": "Lokalna Baza Danych",
		"data":   weather,
	})
}

func (wc *WeatherController) PostWeather(c echo.Context) error {
	req := new(WeatherRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowy format danych"})
	}
	if req.City == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Pole 'city' jest wymagane"})
	}

	var weather Weather
	result := wc.DB.Where("LOWER(city) = ?", strings.ToLower(req.City)).First(&weather)
	
	if result.Error != nil {
		proxyData, err := wc.Proxy.FetchWeather(req.City)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Nie znaleziono w bazie, a zewnętrzne API zwróciło błąd: " + err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"source": "Zewnętrzne API (Proxy)",
			"data":   proxyData,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"source": "Lokalna Baza Danych",
		"data":   weather,
	})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := initDB()
	proxy := &WeatherProxy{}

	weatherController := &WeatherController{
		DB:    db,
		Proxy: proxy,
	}

	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.PostWeather)

	e.Logger.Fatal(e.Start(":8080"))
}