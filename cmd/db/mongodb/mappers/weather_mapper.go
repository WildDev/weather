package mappers

import (
	"app/cmd/db/mongodb/doc"
	"app/cmd/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

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
		ValueC:      src.ValueC,
		MinValueC:   src.MinValueC,
		MaxValueC:   src.MaxValueC,
		ValueF:      src.ValueF,
		MinValueF:   src.MinValueF,
		MaxValueF:   src.MaxValueF,
		Timestamp:   src.Timestamp,
		LastUpdated: src.LastUpdated,
	}, nil
}

func MapModel(src *doc.Weather) *models.Weather {
	return &models.Weather{Id: src.Id.Hex(),
		Country:     src.Country,
		City:        src.City,
		ValueC:      src.ValueC,
		MinValueC:   src.MinValueC,
		MaxValueC:   src.MaxValueC,
		ValueF:      src.ValueF,
		MinValueF:   src.MinValueF,
		MaxValueF:   src.MaxValueF,
		Timestamp:   src.Timestamp,
		LastUpdated: src.LastUpdated,
	}
}
