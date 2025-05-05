package db

import "app/cmd/models"

type WeatherDao interface {
	Add(item models.Weather) (*models.Weather, error)
	FindTopByCountryAndCityOrderByTimestamp(country string, city string) (*models.Weather, error)
}
