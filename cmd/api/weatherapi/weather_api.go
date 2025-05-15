package weatherapi

import (
	"app/cmd"
	"app/cmd/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Location struct {
	Country string `json:"country"`
	City    string `json:"name"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC            float32    `json:"temp_c"`
	TempF            float32    `json:"temp_f"`
	LastUpdatedEpoch int64      `json:"last_updated_epoch"`
	Condition        *Condition `json:"condition"`
}

type ForecastRecord struct {
	ForecastDay []ForecastDayItem `json:"forecastday"`
}

type Forecast struct {
	Location *Location       `json:"location"`
	Current  *Current        `json:"current"`
	Forecast *ForecastRecord `json:"forecast"`
}

type ForecastDay struct {
	MinTempC float64 `json:"mintemp_c"`
	MinTempF float64 `json:"mintemp_f"`
	MaxTempC float64 `json:"maxtemp_c"`
	MaxTempF float64 `json:"maxtemp_f"`
}

type ForecastDayItem struct {
	Day *ForecastDay `json:"day"`
}

type WeatherApi struct {
	Config *cmd.Api
}

func (w *WeatherApi) Init() {

}

func (w *WeatherApi) Destroy() {

}

func (c *Condition) String() string {
	return fmt.Sprintf("Text=%s", c.Text)
}

func (c *Current) String() string {
	return fmt.Sprintf("TempC=%v TempF=%v LastUpdatedEpoch=%v Condition=(%v)", c.TempC, c.TempF, c.LastUpdatedEpoch, c.Condition)
}

func (f *ForecastRecord) String() string {
	return fmt.Sprintf("ForecastDay=%v", f.ForecastDay)
}

func (f *Forecast) String() string {
	return fmt.Sprintf("Current=(%v) Forecast=(%v)", f.Current, f.Forecast)
}

func (f *ForecastDay) String() string {
	return fmt.Sprintf("MinTempC=%v MinTempF=%v MaxTempC=%v MaxTempF=%v", f.MinTempC, f.MinTempF, f.MaxTempC, f.MaxTempF)
}

func (f *ForecastDayItem) String() string {
	return fmt.Sprintf("Day=(%v)", f.Day)
}

func (api *WeatherApi) buildQueryString(country string, city string) string {
	return fmt.Sprintf("%s,%s", city, country)
}

func (api *WeatherApi) buildRequest(country string, city string) (*http.Request, error) {

	config := api.Config

	if req, err := http.NewRequest(http.MethodGet, config.Url, nil); err == nil {

		q := req.URL.Query()

		q.Add("key", config.SecretKey)
		q.Add("q", api.buildQueryString(country, city))

		req.URL.RawQuery = q.Encode()

		return req, nil
	} else {
		return nil, err
	}
}

func MapModel(src *Forecast) *models.Weather {

	l := src.Location
	c := src.Current
	f := src.Forecast.ForecastDay[0].Day

	timestamp := time.Now()
	lastUpdated := cmd.EpochToTime(c.LastUpdatedEpoch)

	return &models.Weather{
		Country:     l.Country,
		City:        l.City,
		ValueC:      int(c.TempC),
		MinValueC:   int(f.MinTempC),
		MaxValueC:   int(f.MaxTempC),
		ValueF:      int(c.TempF),
		MinValueF:   int(f.MinTempF),
		MaxValueF:   int(f.MaxTempF),
		Timestamp:   &timestamp,
		LastUpdated: &lastUpdated,
	}
}

func (api *WeatherApi) GetForecast(country string, city string) (*models.Weather, error) {

	if req_str, req_str_err := api.buildRequest(country, city); req_str_err == nil {

		if req, req_err := (&http.Client{}).Do(req_str); req_err == nil {

			var forecast Forecast

			r_data, r_data_err := io.ReadAll(req.Body)
			defer req.Body.Close()

			if r_data_err == nil {

				if json_err := json.Unmarshal(r_data, &forecast); json_err == nil {
					return MapModel(&forecast), nil
				} else {
					return nil, json_err
				}

			} else {
				return nil, r_data_err
			}

		} else {
			return nil, req_err
		}
	} else {
		return nil, req_str_err
	}
}
