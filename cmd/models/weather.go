package models

import (
	"fmt"
	"time"
)

type Weather struct {
	Id          string
	Country     string
	City        string
	Now         *WeatherNow
	Today       *WeatherForecastDay
	Forecast    []*WeatherForecast
	Timestamp   *time.Time
	LastUpdated *time.Time
	Stale       bool
}

type WeatherNow struct {
	ValueC    int
	ValueF    int
	Condition string
}

type WeatherForecast struct {
	Date      string
	DateEpoch int64
	Day       *WeatherForecastDay
}

type WeatherForecastDay struct {
	MinValueC int
	MaxValueC int
	MinValueF int
	MaxValueF int
	Condition string
}

func (w *Weather) String() string {
	return fmt.Sprintf("Id=%s Country=%s City=%s Now=%v Today=%v Forecast=%v Timestamp=%v LastUpdated=%v Stale=%v", w.Id, w.Country, w.City, *w.Now, *w.Today, w.Forecast, *w.Timestamp, *w.LastUpdated, w.Stale)
}

func (w *WeatherNow) String() string {
	return fmt.Sprintf("ValueC=%d ValueF=%d Condition=%s", w.ValueC, w.ValueF, w.Condition)
}

func (w *WeatherForecast) String() string {
	return fmt.Sprintf("Date=%s DateEpoch=%d Day=%v", w.Date, w.DateEpoch, *w.Day)
}

func (w *WeatherForecastDay) String() string {
	return fmt.Sprintf("MinValueC=%d MaxValueC=%d MinValueF=%d MaxValueF=%d Condition=%s", w.MinValueC, w.MaxValueC, w.MinValueF, w.MaxValueF, w.Condition)
}
