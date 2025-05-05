package main

import (
	"app/cmd"
	"app/cmd/api/weatherapi"
	"app/cmd/db/mongodb"
	"app/cmd/db/mongodb/dao"
	"app/cmd/rest"
	"app/cmd/services"
	"log"

	"github.com/gin-gonic/gin"
)

func initMongoConn(config *cmd.Config) *mongodb.MongoConn {

	if cfg := config.MongoDB; cfg == nil {

		log.Fatalln("MongoDB connection details required")
		return nil

	} else {
		return &mongodb.MongoConn{Config: cfg}
	}
}

func configureAndStart(config *cmd.Config, endpoint *rest.WeatherEndpoint) {

	host := config.HostBinding
	router := gin.Default()

	router.GET("/now", endpoint.Now)

	if tls := config.TLS; tls.Use {
		log.Println("TLS enabled")
		router.RunTLS(host, tls.CertPath, tls.KeyPath)
	} else {
		log.Println("TLS disabled")
		router.Run(host)
	}
}

func main() {

	config := cmd.ReadEnv()
	log.Println(config.String())

	ctx := cmd.Context{}
	defer ctx.DestroyAll()

	conn := initMongoConn(&config)
	dao := &dao.WeatherDao{Conn: conn}
	api := &weatherapi.WeatherApi{Config: config.Api}
	service := &services.WeatherService{Config: &config, Api: api, Dao: dao}
	endpoint := &rest.WeatherEndpoint{Service: service}

	ctx.Add(conn)
	ctx.Add(dao)
	ctx.Add(api)
	ctx.Add(service)
	ctx.Add(endpoint)

	configureAndStart(&config, endpoint)
}
