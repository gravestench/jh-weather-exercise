package openWeatherMap

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/jh-weather-exercise/pkg/ginWebServer"
)

var _ ginWebServer.GinRouteInitializer = &Service{}

func (s *Service) InitializeGinRoutes(g *gin.RouterGroup) {
	g.GET("weather/current", s.handleGetCurrentWeather)
	g.GET("weather/current/describe", s.handleDescribeCurrentWeather)
}

func (s *Service) handleGetCurrentWeather(c *gin.Context) {
	type request struct {
		Latitude, Longitude float64
	}

	var params request

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("parsing json request: %v", err))
		return
	}

	res, err := s.GetCurrentWeatherData(params.Latitude, params.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("getting current weather: %v", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Service) handleDescribeCurrentWeather(c *gin.Context) {
	type request struct {
		Latitude, Longitude float64
	}

	var params request

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("parsing json request: %v", err))
		return
	}

	res, err := s.GetCurrentWeatherData(params.Latitude, params.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("getting current weather: %v", err))
		return
	}

	if res == nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("got nil weather response"))
		return
	}

	type payload struct {
		Description string
	}

	c.JSON(http.StatusOK, payload{Description: DescribeWeather(*res)})
}
