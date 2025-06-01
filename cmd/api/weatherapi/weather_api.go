package weatherapi

import (
	"app/cmd"
	"app/cmd/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	Location *WeatherLocation     `json:"location"`
	Current  *WeatherNow          `json:"current"`
	Forecast *WeatherForecastNode `json:"forecast"`
}

type WeatherLocation struct {
	Country string `json:"country"`
	Region  string `json:"region"`
}

type WeatherCondition struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type WeatherNow struct {
	TempC            float32           `json:"temp_c"`
	TempF            float32           `json:"temp_f"`
	LastUpdatedEpoch int64             `json:"last_updated_epoch"`
	Condition        *WeatherCondition `json:"condition"`
}

type WeatherForecastDay struct {
	MinTempC  float64           `json:"mintemp_c"`
	MaxTempC  float64           `json:"maxtemp_c"`
	MinTempF  float64           `json:"mintemp_f"`
	MaxTempF  float64           `json:"maxtemp_f"`
	Condition *WeatherCondition `json:"condition"`
}

type WeatherForecast struct {
	Date      string              `json:"date"`
	DateEpoch int64               `json:"date_epoch"`
	Day       *WeatherForecastDay `json:"day"`
}

type WeatherForecastNode struct {
	ForecastDay []*WeatherForecast `json:"forecastday"`
}

type WeatherApi struct {
	Config *cmd.Api
}

func (w *WeatherApi) Init() {

}

func (w *WeatherApi) Destroy() {

}

func (w *Weather) String() string {
	return fmt.Sprintf("Location=%v Current=%v Forecast=%v", *w.Location, *w.Current, *w.Forecast)
}

func (w *WeatherLocation) String() string {
	return fmt.Sprintf("Country=%s Region=%s", w.Country, w.Region)
}

func (w *WeatherCondition) String() string {
	return fmt.Sprintf("Code=%d Text=%s", w.Code, w.Text)
}

func (w *WeatherNow) String() string {
	return fmt.Sprintf("TempC=%f TempF=%f LastUpdatedEpoch=%v Condition=%v", w.TempC, w.TempF, w.LastUpdatedEpoch, *w.Condition)
}

func (w *WeatherForecastDay) String() string {
	return fmt.Sprintf("MinTempC=%f MaxTempC=%f MinTempF=%f MaxTempF=%f Condition=%v", w.MinTempC, w.MinTempF, w.MaxTempC, w.MaxTempF, *w.Condition)
}

func (w *WeatherForecast) String() string {
	return fmt.Sprintf("Date=%s DateEpoch=%d Day=%v", w.Date, w.DateEpoch, *w.Day)
}

func (w *WeatherForecastNode) String() string {
	return fmt.Sprintf("ForecastDay=%v", w.ForecastDay)
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

func mapConditionCode(code int) string {

	switch code {

	case 1000:
		return "clear"
	case 1003, 1006:
		return "partial-clear"
	case 1009:
		return "clouds"
	case 1030, 1135:
		return "fog"
	case 1150, 1153:
		return "drizzle"
	case 1072, 1147, 1168:
		return "freezing-drizzle"
	case 1171:
		return "heavy-freezing-drizzle"
	case 1063, 1180, 1186, 1192:
		return "partial-rain"
	case 1183, 1189, 1195:
		return "rain"
	case 1240, 1243, 1246:
		return "shower"
	case 1087:
		return "storm"
	case 1273, 1276:
		return "rain-storm"
	case 1069, 1198, 1204, 1237, 1249, 1261:
		return "freezing-rain"
	case 1201, 1207, 1252, 1264:
		return "heavy-freezing-rain"
	case 1066, 1210, 1216, 1222:
		return "partial-snow"
	case 1213, 1219, 1225, 1255:
		return "snow"
	case 1114, 1117, 1258:
		return "blizzard"
	case 1279, 1282:
		return "snow-storm"
	default:
		log.Println("Unmapped condition code", code)
		return ""
	}
}

func mapWeatherNow(src *WeatherNow) *models.WeatherNow {
	return &models.WeatherNow{
		ValueC:    int(src.TempC),
		ValueF:    int(src.TempF),
		Condition: mapConditionCode(src.Condition.Code),
	}
}

func mapWeatherForecastDay(src *WeatherForecastDay) *models.WeatherForecastDay {
	return &models.WeatherForecastDay{
		MinValueC: int(src.MinTempC),
		MaxValueC: int(src.MaxTempC),
		MinValueF: int(src.MinTempF),
		MaxValueF: int(src.MaxTempF),
		Condition: mapConditionCode(src.Condition.Code),
	}
}

func mapWeatherForecast(src *WeatherForecast) *models.WeatherForecast {
	return &models.WeatherForecast{
		Date:      src.Date,
		DateEpoch: src.DateEpoch,
		Day:       mapWeatherForecastDay(src.Day),
	}
}

func mapWeatherForecastArray(src []*WeatherForecast) []*models.WeatherForecast {

	r := make([]*models.WeatherForecast, len(src))

	for i, p := range src {
		r[i] = mapWeatherForecast(p)
	}

	return r
}

func mapWeather(country string, city string, src *Weather) *models.Weather {

	var today *models.WeatherForecastDay = nil
	var forecast []*models.WeatherForecast = nil

	c := src.Current

	timestamp := time.Now()
	lastUpdated := cmd.EpochToTime(c.LastUpdatedEpoch)

	if f := src.Forecast.ForecastDay; len(f) > 0 {

		today = mapWeatherForecastDay(f[0].Day)
		forecast = mapWeatherForecastArray(f[1:])
	} else {
		log.Println("No forecast data found!")
	}

	return &models.Weather{
		Country:     country,
		City:        city,
		Now:         mapWeatherNow(src.Current),
		Today:       today,
		Forecast:    forecast,
		Timestamp:   &timestamp,
		LastUpdated: &lastUpdated,
	}
}

func (api *WeatherApi) GetForecast(country string, city string) (*models.Weather, error) {

	if req_str, req_str_err := api.buildRequest(country, city); req_str_err == nil {

		if req, req_err := (&http.Client{}).Do(req_str); req_err == nil {

			var forecast *Weather

			r_data, r_data_err := io.ReadAll(req.Body)
			defer req.Body.Close()

			if r_data_err == nil {

				if json_err := json.Unmarshal(r_data, &forecast); json_err == nil {
					return mapWeather(country, city, forecast), nil
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
