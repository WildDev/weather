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
		Options: options.Index().SetName("country_1_city_1"),
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

func (dao *WeatherDao) Add(item models.Weather) (*models.Weather, error) {

	if d, d_err := mappers.MapDoc(&item); d_err == nil {

		if r, r_err := dao.Conn.Ref().Collection(collection).InsertOne(context.TODO(), d); r_err == nil {

			item.Id = r.InsertedID.(bson.ObjectID).Hex()
			return &item, nil
		} else {
			return nil, r_err
		}
	} else {
		return nil, d_err
	}
}

func (dao *WeatherDao) FindTopByCountryAndCityOrderByTimestamp(country string, city string) (*models.Weather, error) {

	var item *doc.Weather

	filter := bson.D{
		{Key: "country", Value: country},
		{Key: "city", Value: city},
	}

	opts := options.FindOne().SetSort(bson.D{
		{Key: "timestamp", Value: -1},
	})

	r := dao.Conn.Ref().Collection(collection).FindOne(context.TODO(), filter, opts)

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
