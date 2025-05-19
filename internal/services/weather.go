package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

type WeatherService struct {
	APIKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{APIKey: apiKey}
}

func (s *WeatherService) GetWeather(city string) (WeatherData, error) {
	fmt.Println("Getting weather with API token: ", s.APIKey)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.APIKey, city)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return WeatherData{}, fmt.Errorf("weather API error: status %d", resp.StatusCode)
	}

	var raw struct {
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
		Current struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return WeatherData{}, err
	}

	return WeatherData{
		City:        raw.Location.Name,
		Temperature: raw.Current.TempC,
		Condition:   raw.Current.Condition.Text,
	}, nil
}
