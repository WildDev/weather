FROM golang:1.24

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v
RUN ls

FROM ubuntu:24.10

ENV GIN_MODE='release'

ENV API_SECRET_KEY=''
ENV API_URL='https://api.weatherapi.com/v1/forecast.json'
ENV CACHE_TIMEOUT='20m'
ENV HOST_BINDING=':8080'
ENV LOG_MONGO_URI='false'
ENV MONGO_DATABASE='weather'
ENV MONGO_URI=''
ENV TLS_CERT_PATH='tls/cert.pem'
ENV TLS_KEY_PATH='tls/private.key'
ENV TLS_USE='false'

COPY --from=0 /go/app /

ENTRYPOINT "./app"
