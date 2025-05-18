### Description

Weather checking service. It allows to query weather information based on a user location.

Currently, is only integrated with [weatherapi.com](https://www.weatherapi.com)

[![build](https://github.com/WildDev/weather/actions/workflows/go.yml/badge.svg)](https://github.com/WildDev/weather/actions/workflows/go.yml) [![docker](https://github.com/WildDev/weather/actions/workflows/docker-image.yml/badge.svg)](https://github.com/WildDev/weather/actions/workflows/docker-image.yml)

### How it works

1. Set `country` and `city` request params
3. Send the request

Example:
```bash
curl "https://test.website/now?country=Spain&city=Ibiza" -v

< HTTP/1.1 200
< Content-Type: application/json
<
{
    "value": {

        // Celsius
        "c": {
            "max": 18,
            "min": 13,
            "val": 15
        },

        // Fahrenheit
        "f": {
            "max": 65,
            "min": 60,
            "val": 62
        }
    },
    "stale": false,
    "updated": "2025-05-06T17:15:00Z"
}
```

> [!NOTE]
> If there was no connection between this and 3d-party services for long time, then `stale` flag is set to `true` and the last known weather value is returned.

### Get started

How to build and run:

```bash
export MONGO_URI="mongodb+srv://weather:test@mongodb.example.com/?tls=true&authSource=admin&replicaSet=mongodb"

export API_SECRET_KEY="586e596cb3ebe58d4e6913fd5df2"

cd weather
go build

./app
```

Environment variables and defaults:

```bash
API_SECRET_KEY=''
API_URL='https://api.weatherapi.com/v1/forecast.json'
CACHE_TIMEOUT='20m'
HOST_BINDING=':8080'
LOG_MONGO_URI='false'
MONGO_DATABASE='weather'
MONGO_URI=''
TLS_CERT_PATH='tls/cert.pem'
TLS_KEY_PATH='tls/private.key'
TLS_USE='false'
```

Also available on [Docker Hub](https://hub.docker.com/r/wilddev/weather)
