package endpoints

import (
	"app/cmd/rest"
	"app/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherEndpoint struct {
	Service *services.WeatherService
}

const country_param, city_param = "country", "city"

func (w *WeatherEndpoint) Init() {

}

func (w *WeatherEndpoint) Destroy() {

}

func (e *WeatherEndpoint) Now(c *gin.Context) {

	const country_param, city_param = "country", "city"
	var country, city string

	if country = c.Query(country_param); country == "" {
		c.JSON(http.StatusBadRequest, rest.FieldError{
			Field: country_param,
			Error: "country is not set",
		})
		return
	}

	if city = c.Query(city_param); city == "" {
		c.JSON(http.StatusBadRequest, rest.FieldError{
			Field: city_param,
			Error: "city is not set",
		})
		return
	}

	if w, w_err := e.Service.Now(country, city); w_err == nil {
		c.JSON(http.StatusOK, gin.H{
			"ValueC":      w.ValueC,
			"MinValueC":   w.MinValueC,
			"MaxValueC":   w.MaxValueC,
			"ValueF":      w.ValueF,
			"MinValueF":   w.MinValueF,
			"MaxValueF":   w.MaxValueF,
			"LastUpdated": w.LastUpdated,
			"Stale":       w.Stale,
		})
	} else {

		c.JSON(http.StatusInternalServerError, rest.GlobalError{Error: "something went wrong"})

		panic(w_err)
	}
}
