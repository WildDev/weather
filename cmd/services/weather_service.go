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

func scaleMinMax(item *models.Weather) {

	value_c := item.Now.ValueC
	value_f := item.Now.ValueF

	today := item.Today

	if value_c < today.MinValueC {
		today.MinValueC = value_c
	}

	if value_c > today.MaxValueC {
		today.MaxValueC = value_c
	}

	if value_f < today.MinValueF {
		today.MinValueF = value_f
	}

	if value_f > today.MaxValueF {
		today.MaxValueF = value_f
	}
}

func (service *WeatherService) Now(country string, city string) (*models.Weather, error) {

	if r, r_err := service.Dao.FindByCountryAndCity(country, city); r_err == nil {

		now := time.Now().UTC()
		b, b_err := cmd.Forward(&now, service.Config.CacheTimeout)

		if b_err == nil {
			if r == nil || r.Timestamp.Before(*b) {

				if f, f_err := service.Api.GetForecast(country, city); f_err == nil {

					if r != nil && r.LastUpdated.Compare(*f.LastUpdated) >= 0 {
						return r, nil
					} else {

						scaleMinMax(f)

						if k, k_err := service.Dao.Upsert(f); k_err == nil {
							return k, nil
						} else {
							return nil, k_err
						}
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
