package models

import (
	"fmt"
	"time"
)

type Weather struct {
	Id          string
	Country     string
	City        string
	ValueC      int
	MinValueC   int
	MaxValueC   int
	ValueF      int
	MinValueF   int
	MaxValueF   int
	Timestamp   *time.Time
	LastUpdated *time.Time
	Stale       bool
}

func (w *Weather) String() string {

	return fmt.Sprintf("Id=%s Country=%s City=%s ValueC=%d MinValueC=%d MaxValueC=%d ValueF=%d MinValueF=%d MaxValueF=%d Timestamp=%v LastUpdated=%v Stale=%v",
		w.Id, w.Country, w.City, w.ValueC, w.MinValueC, w.MaxValueC, w.ValueF, w.MinValueF, w.MaxValueF, *w.Timestamp, *w.LastUpdated, w.Stale)
}
