package db

import "app/cmd/models"

type WeatherDao interface {
	Upsert(item *models.Weather) (*models.Weather, error)
	FindByCountryAndCity(country string, city string) (*models.Weather, error)
}
