package mongodb

import (
	"app/cmd"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConn struct {
	Config *cmd.MongoDB
	Client *mongo.Client
}

func (c *MongoConn) Init() {
	c.Connect()
}

func (c *MongoConn) Destroy() {
	c.Disconnect()
}

func (c *MongoConn) Ref() *mongo.Database {
	return c.Client.Database(c.Config.Database)
}

func (c *MongoConn) Connect() {

	config := c.Config

	if client, err := mongo.Connect(options.Client().ApplyURI(config.Uri)); err == nil {

		c.Client = client
		log.Println("MongoDB connection ready")

		return

	} else {
		log.Fatal("Failed to connect", err)
	}
}

func (c *MongoConn) Disconnect() {

	if c == nil {

		log.Println("No active connection")
		return
	}

	c.Client.Disconnect(context.TODO())
	c.Client = nil

	log.Println("Disconnected")
}
