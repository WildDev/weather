package doc

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Weather struct {
	Id          *bson.ObjectID      `bson:"_id,omitempty"`
	Country     string              `bson:"country,omitempty"`
	City        string              `bson:"city,omitempty"`
	Now         *WeatherNow         `bson:"now,omitempty"`
	Today       *WeatherForecastDay `bson:"today,omitempty"`
	Forecast    []*WeatherForecast  `bson:"forecast,omitempty"`
	Timestamp   *time.Time          `bson:"timestamp,omitempty"`
	LastUpdated *time.Time          `bson:"last_updated,omitempty"`
}

type WeatherNow struct {
	ValueC    int    `bson:"value_c,omitempty"`
	ValueF    int    `bson:"value_f,omitempty"`
	Condition string `bson:"condition,omitempty"`
}

type WeatherForecast struct {
	Date string              `bson:"date,omitempty"`
	Day  *WeatherForecastDay `bson:"day,omitempty"`
}

type WeatherForecastDay struct {
	MinValueC int    `bson:"min_value_c,omitempty"`
	MaxValueC int    `bson:"max_value_c,omitempty"`
	MinValueF int    `bson:"min_value_f,omitempty"`
	MaxValueF int    `bson:"max_value_f,omitempty"`
	Condition string `bson:"condition,omitempty"`
}

func (w *Weather) GetIdAsString() string {

	if w.Id == nil {
		return ""
	} else {
		return w.Id.Hex()
	}
}

func (w *Weather) String() string {
	return fmt.Sprintf("Id=%s Country=%s City=%s Now=%v Today=%v Forecast=%v Timestamp=%v LastUpdated=%v",
		w.GetIdAsString(), w.Country, w.City, *w.Now, *w.Today, w.Forecast, *w.Timestamp, *w.LastUpdated)
}

func (w *WeatherNow) String() string {
	return fmt.Sprintf("ValueC=%d ValueF=%d Condition=%s", w.ValueC, w.ValueF, w.Condition)
}

func (w *WeatherForecast) String() string {
	return fmt.Sprintf("Date=%s Day=%v", w.Date, *w.Day)
}

func (w *WeatherForecastDay) String() string {
	return fmt.Sprintf("MinValueC=%d MaxValueC=%d MinValueF=%d MaxValueF=%d Condition=%s", w.MinValueC, w.MaxValueC, w.MinValueF, w.MaxValueF, w.Condition)
}
