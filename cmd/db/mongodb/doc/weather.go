package doc

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Weather struct {
	Id          *bson.ObjectID `bson:"_id,omitempty"`
	Country     string         `bson:"country,omitempty"`
	City        string         `bson:"city,omitempty"`
	ValueC      int            `bson:"value_c,omitempty"`
	MinValueC   int            `bson:"min_value_c,omitempty"`
	MaxValueC   int            `bson:"max_value_c,omitempty"`
	ValueF      int            `bson:"value_f,omitempty"`
	MinValueF   int            `bson:"min_value_f,omitempty"`
	MaxValueF   int            `bson:"max_value_f,omitempty"`
	Timestamp   *time.Time     `bson:"timestamp,omitempty"`
	LastUpdated *time.Time     `bson:"last_updated,omitempty"`
}

func (w *Weather) GetIdAsString() string {

	if w.Id == nil {
		return ""
	} else {
		return w.Id.Hex()
	}
}

func (w *Weather) String() string {
	return fmt.Sprintf("Id=%s Country=%s City=%s ValueC=%d MinValueC=%d MaxValueC=%d ValueF=%d MinValueF=%d MaxValueF=%d Timestamp=%v LastUpdated=%v",
		w.GetIdAsString(), w.Country, w.City, w.ValueC, w.MinValueC, w.MaxValueC, w.ValueF, w.MinValueF, w.MaxValueF, *w.Timestamp, *w.LastUpdated)
}
