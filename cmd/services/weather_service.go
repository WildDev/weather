package services

import (
	"app/cmd"
	"app/cmd/api"
	"app/cmd/db"
	"app/cmd/models"
	"log"
	"time"
)

type WeatherService struct {
	Config *cmd.Config
	Api    api.WeatherApi
	Dao    db.WeatherDao
}

func (service *WeatherService) Init() {

}

func (service *WeatherService) Destroy() {

}

func (service *WeatherService) Now(country string, city string) (*models.Weather, error) {

	if r, r_err := service.Dao.FindTopByCountryAndCityOrderByTimestamp(country, city); r_err == nil {

		now := time.Now().UTC()
		b, b_err := cmd.Forward(&now, service.Config.CacheTimeout)

		if b_err == nil {
			if r == nil || r.Timestamp.Before(*b) {

				if f, f_err := service.Api.GetForecast(country, city); f_err == nil {

					if r != nil && r.LastUpdated.Compare(*f.LastUpdated) >= 0 {
						return r, nil
					} else if k, k_err := service.Dao.Add(*f); k_err == nil {
						return k, nil
					} else {
						return nil, k_err
					}

				} else {

					if r != nil {

						r.Stale = true
						log.Println("Failed to query the forecast", f_err)

						return r, nil
					}

					return nil, f_err
				}
			} else {
				return r, nil
			}
		} else {
			return nil, b_err
		}
	} else {
		return nil, r_err
	}
}
