package dao

import (
	"app/cmd/db/mongodb"
	"app/cmd/db/mongodb/doc"
	"app/cmd/db/mongodb/mappers"
	"app/cmd/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const collection = "weather"

var indexes = []mongo.IndexModel{
	{
		Keys: bson.D{
			{Key: "country", Value: 1},
			{Key: "city", Value: 1},
		},
		Options: options.Index().SetName("country_1_city_1").SetUnique(true),
	},
}

type WeatherDao struct {
	Conn *mongodb.MongoConn
}

func (dao *WeatherDao) Init() {
	dao.createIndexes()
}

func (dao *WeatherDao) Destroy() {

}

func (dao *WeatherDao) createIndexes() {
	for i := range indexes {
		if r, r_err := dao.Conn.Ref().Collection(collection).Indexes().CreateOne(context.TODO(), indexes[i]); r_err == nil {
			log.Printf("Index added or updated: %s.%s\n", collection, r)
		} else {
			panic(r_err)
		}
	}
}

func createCountryAndCityFilter(country string, city string) *bson.D {
	return &bson.D{
		{Key: "country", Value: country},
		{Key: "city", Value: city},
	}
}

func (dao *WeatherDao) Upsert(item *models.Weather) (*models.Weather, error) {

	if d, d_err := mappers.MapDoc(item); d_err == nil {

		opts := options.Replace().SetUpsert(true)

		if r, r_err := dao.Conn.Ref().Collection(collection).ReplaceOne(context.TODO(), createCountryAndCityFilter(item.Country, item.City), d, opts); r_err == nil {

			id := r.UpsertedID

			if id != nil {
				item.Id = id.(bson.ObjectID).Hex()
			}

			return item, nil
		} else {
			return nil, r_err
		}
	} else {
		return nil, d_err
	}
}

func (dao *WeatherDao) FindByCountryAndCity(country string, city string) (*models.Weather, error) {

	var item *doc.Weather

	r := dao.Conn.Ref().Collection(collection).FindOne(context.TODO(),
		createCountryAndCityFilter(country, city))

	if err := r.Decode(&item); err == nil {
		return mappers.MapModel(item), nil
	} else {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}
}
