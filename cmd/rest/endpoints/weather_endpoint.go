package endpoints

import (
	"app/cmd/rest"
	"app/cmd/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherEndpoint struct {
	Service *services.WeatherService
}

func (e *WeatherEndpoint) Init() {

}

func (e *WeatherEndpoint) Destroy() {

}

func (e *WeatherEndpoint) reportInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, rest.GlobalError{Error: "something went wrong"})
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

		now := w.Now
		today := w.Today

		if today == nil {

			log.Println("No forecast data found for today!")
			e.reportInternalServerError(c)

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"value": gin.H{
				"c": gin.H{
					"val": now.ValueC,
					"min": today.MinValueC,
					"max": today.MaxValueC,
				},
				"f": gin.H{
					"val": now.ValueF,
					"min": today.MinValueF,
					"max": today.MaxValueF,
				},
			},
			"condition": now.Condition,
			"updated":   w.LastUpdated,
			"stale":     w.Stale,
		})
	} else {

		e.reportInternalServerError(c)
		panic(w_err)
	}
}
