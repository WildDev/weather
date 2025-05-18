package mappers

import (
	"app/cmd/db/mongodb/doc"
	"app/cmd/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func mapDocWeatherNow(src *models.WeatherNow) *doc.WeatherNow {
	return &doc.WeatherNow{
		ValueC:    src.ValueC,
		ValueF:    src.ValueF,
		Condition: src.Condition,
	}
}

func mapDocWeatherForecastDay(src *models.WeatherForecastDay) *doc.WeatherForecastDay {
	return &doc.WeatherForecastDay{
		MinValueC: src.MinValueC,
		MaxValueC: src.MaxValueC,
		MinValueF: src.MinValueF,
		MaxValueF: src.MaxValueF,
		Condition: src.Condition,
	}
}

func mapDocWeatherForecast(src *models.WeatherForecast) *doc.WeatherForecast {
	return &doc.WeatherForecast{
		Date: src.Date,
		Day:  mapDocWeatherForecastDay(src.Day),
	}
}

func mapDocWeatherForecastArray(src []*models.WeatherForecast) []*doc.WeatherForecast {

	r := make([]*doc.WeatherForecast, len(src))

	for i, p := range src {
		r[i] = mapDocWeatherForecast(p)
	}

	return r
}

func MapDoc(src *models.Weather) (*doc.Weather, error) {

	var _id_p *bson.ObjectID

	if id := src.Id; id != "" {

		_id, _id_err := bson.ObjectIDFromHex(id)

		if _id_err == nil {
			_id_p = &_id
		} else {
			return nil, _id_err
		}
	}

	return &doc.Weather{Id: _id_p,
		Country:     src.Country,
		City:        src.City,
		Now:         mapDocWeatherNow(src.Now),
		Forecast:    mapDocWeatherForecastArray(src.Forecast),
		Timestamp:   src.Timestamp,
		LastUpdated: src.LastUpdated,
	}, nil
}

func mapModelWeatherNow(src *doc.WeatherNow) *models.WeatherNow {
	return &models.WeatherNow{
		ValueC:    src.ValueC,
		ValueF:    src.ValueF,
		Condition: src.Condition,
	}
}

func mapModelWeatherForecastDay(src *doc.WeatherForecastDay) *models.WeatherForecastDay {
	return &models.WeatherForecastDay{
		MinValueC: src.MinValueC,
		MaxValueC: src.MaxValueC,
		MinValueF: src.MinValueF,
		MaxValueF: src.MaxValueF,
		Condition: src.Condition,
	}
}

func mapModelWeatherForecast(src *doc.WeatherForecast) *models.WeatherForecast {
	return &models.WeatherForecast{
		Date: src.Date,
		Day:  mapModelWeatherForecastDay(src.Day),
	}
}

func mapModelWeatherForecastArray(src []*doc.WeatherForecast) []*models.WeatherForecast {

	r := make([]*models.WeatherForecast, len(src))

	for i, p := range src {
		r[i] = mapModelWeatherForecast(p)
	}

	return r
}

func MapModel(src *doc.Weather) *models.Weather {
	return &models.Weather{Id: src.Id.Hex(),
		Country:     src.Country,
		City:        src.City,
		Now:         mapModelWeatherNow(src.Now),
		Forecast:    mapModelWeatherForecastArray(src.Forecast),
		Timestamp:   src.Timestamp,
		LastUpdated: src.LastUpdated,
	}
}
