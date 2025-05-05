package api

import "app/cmd/models"

type WeatherApi interface {
	GetForecast(city string) (*models.Weather, error)
}
