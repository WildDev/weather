package cmd

import (
	"fmt"
	"log"
	"os"
)

const API_SECRET_KEY, API_URL, CACHE_TIMEOUT, HOST_BINDING, LOG_MONGO_URI, MONGO_DATABASE, MONGO_URI, TLS_CERT_PATH, TLS_KEY_PATH, TLS_USE = "API_SECRET_KEY", "API_URL", "CACHE_TIMEOUT", "HOST_BINDING", "LOG_MONGO_URI", "MONGO_DATABASE", "MONGO_URI", "TLS_CERT_PATH", "TLS_KEY_PATH", "TLS_USE"

const masked_text = "*masked*"

var defaults = map[string]string{
	HOST_BINDING:   ":8080",
	CACHE_TIMEOUT:  "20m",
	LOG_MONGO_URI:  "false",
	MONGO_DATABASE: "weather",
	TLS_USE:        "false",
	TLS_CERT_PATH:  "tls/cert.pem",
	TLS_KEY_PATH:   "tls/private.key",
}

type TLS struct {
	Use      bool
	CertPath string
	KeyPath  string
}

type MongoDB struct {
	Uri      string
	LogUri   string
	Database string
}

type Api struct {
	Url       string
	SecretKey string
}

type Config struct {
	TLS          *TLS
	MongoDB      *MongoDB
	Api          *Api
	CacheTimeout string
	HostBinding  string
}

func (t *TLS) String() string {
	return fmt.Sprintf("Use=%v CertPath=%s KeyPath=%s", t.Use, t.CertPath, t.KeyPath)
}

func (c *MongoDB) String() string {

	var uri string

	if c.LogUri == "true" {
		uri = c.Uri
	} else {
		uri = masked_text
	}

	return fmt.Sprintf("Uri=%s Database=%s", uri, c.Database)
}

func (a *Api) String() string {
	return fmt.Sprintf("Url=%s SecretKey=%s", a.Url, masked_text)
}

func (c *Config) String() string {
	return fmt.Sprintf(`

*** Settings dump ***
TLS=(%v)
MongoDB=(%v)
Api=(%v)
CacheTimeout=(%s)
HostBinding=(%s)
---
`, c.TLS, c.MongoDB, c.Api, c.CacheTimeout, c.HostBinding)
}

func getVal(key string) string {

	if val := os.Getenv(key); val == "" {
		return defaults[key]
	} else {
		return val
	}
}

func readTLS() *TLS {

	if should_use := getVal(TLS_USE); should_use == "true" {

		var cert_path, key_path string

		if cert_path = getVal(TLS_CERT_PATH); cert_path == "" {

			log.Println("TLS certificate path is required")
			return nil
		}

		if key_path = getVal(TLS_KEY_PATH); key_path == "" {

			log.Println("TLS private key path is required")
			return nil
		}

		return &TLS{Use: true, CertPath: cert_path, KeyPath: key_path}
	} else {
		return &TLS{Use: false}
	}
}

func readMongoDB() *MongoDB {

	uri := getVal(MONGO_URI)
	database := getVal(MONGO_DATABASE)

	if uri == "" {

		log.Println("No MongoDB settings set")
		return nil
	}

	if database == "" {

		log.Fatalln(MONGO_DATABASE, "is required")
		return nil
	}

	return &MongoDB{Uri: uri, LogUri: getVal(LOG_MONGO_URI), Database: database}
}

func readApi() *Api {

	url := getVal(API_URL)
	secret_key := getVal(API_SECRET_KEY)

	if url == "" {

		log.Println("No API settings set")
		return nil
	}

	if secret_key == "" {

		log.Fatalln(API_SECRET_KEY, "is required")
		return nil
	}

	return &Api{Url: url, SecretKey: secret_key}
}

func ReadEnv() Config {

	var tls = readTLS()
	var mongo_db = readMongoDB()
	var api = readApi()

	return Config{
		TLS:          tls,
		MongoDB:      mongo_db,
		Api:          api,
		CacheTimeout: getVal(CACHE_TIMEOUT),
		HostBinding:  getVal(HOST_BINDING),
	}
}
