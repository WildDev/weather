package rest

import (
	"app/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherEndpoint struct {
	Service *services.WeatherService
}

func (w *WeatherEndpoint) Init() {

}

func (w *WeatherEndpoint) Destroy() {

}

func (e *WeatherEndpoint) Now(c *gin.Context) {

	var country, city string

	if country = c.Query("country"); country == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "country is not set",
		})
		return
	}

	if city = c.Query("city"); city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "city is not set",
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

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})

		panic(w_err)
	}
}
