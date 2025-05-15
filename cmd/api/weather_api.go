package api

import "app/cmd/models"

type WeatherApi interface {
	GetForecast(country string, city string) (*models.Weather, error)
}
