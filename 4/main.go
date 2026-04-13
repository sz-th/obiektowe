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
	City   string   `json:"city"` 
	Cities []string `json:"cities"`
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

func (wc *WeatherController) processCity(city string) map[string]interface{} {
	city = strings.TrimSpace(city) 
	
	var weather Weather
	result := wc.DB.Where("LOWER(city) = ?", strings.ToLower(city)).First(&weather)
	
	if result.Error != nil {
		proxyData, err := wc.Proxy.FetchWeather(city)
		if err != nil {
			return map[string]interface{}{
				"city":   city,
				"status": "error",
				"error":  err.Error(),
			}
		}

		if err := wc.DB.Create(proxyData).Error; err != nil {
			return map[string]interface{}{
				"city":   city,
				"status": "error",
				"error":  "Błąd zapisu do bazy danych",
			}
		}

		return map[string]interface{}{
			"city":   city,
			"status": "success",
			"source": "Zewnętrzne API (Zapisano do bazy)",
			"data":   proxyData,
		}
	}

	return map[string]interface{}{
		"city":   city,
		"status": "success",
		"source": "Lokalna Baza Danych",
		"data":   weather,
	}
}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	cityParam := c.QueryParam("city")
	if cityParam == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Podaj miasta w parametrze (po przecinku), np. /weather?city=Krakow,Berlin",
		})
	}

	cities := strings.Split(cityParam, ",")
	var results []map[string]interface{}

	for _, city := range cities {
		if city != "" {
			results = append(results, wc.processCity(city))
		}
	}

	return c.JSON(http.StatusOK, results)
}

func (wc *WeatherController) PostWeather(c echo.Context) error {
	req := new(WeatherRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowy format danych"})
	}

	var targets []string
	if req.City != "" {
		targets = append(targets, req.City)
	}
	targets = append(targets, req.Cities...)

	if len(targets) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Podaj pole 'city' lub tablicę 'cities' w formacie JSON",
		})
	}

	var results []map[string]interface{}
	for _, city := range targets {
		if city != "" {
			results = append(results, wc.processCity(city))
		}
	}

	return c.JSON(http.StatusOK, results)
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