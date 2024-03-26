package openWeatherMap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gravestench/jh-weather-exercise/pkg/openWeatherMap/api"
)

const (
	apiRoot = "https://api.openweathermap.org"
)

type Service struct {
	Config
}

func (s *Service) GetCurrentWeatherData(lat, long float64) (*api.ResponseCurrentWeather, error) {
	url := fmt.Sprintf(api.UrlCurrentWeatherQuery, lat, long, s.ApiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("issuing request: %v", err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	var result api.ResponseCurrentWeather

	if err = json.Unmarshal(responseData, &result); err != nil {
		return nil, fmt.Errorf("unmarshalling response: %v", err)
	}

	return &result, nil
}

// DescribeWeather will take a ResponseCurrentWeather from the
// Open Weather Map API and return a string description.
func DescribeWeather(weather api.ResponseCurrentWeather) (description string) {
	if temp := weather.Main.Temp - 273.15; temp < 0 {
		description += "It's freezing"
	} else if temp < 10 {
		description += "It's cold"
	} else if temp < 20 {
		description += "It's cool"
	} else if temp < 30 {
		description += "It's warm"
	} else {
		description += "It's hot"
	}

	if feelsLike := weather.Main.FeelsLike - 273.15; feelsLike < weather.Main.Temp-273.15 {
		description += ", feels colder than it actually is"
	} else if feelsLike > weather.Main.Temp-273.15 {
		description += ", feels warmer than it actually is"
	}

	if humidity := weather.Main.Humidity; humidity < 30 {
		description += ", and dry"
	} else if humidity < 60 {
		description += ", and moderately humid"
	} else {
		description += ", and humid"
	}

	if clouds := weather.Clouds.All; clouds < 30 {
		description += ", with clear skies"
	} else if clouds < 60 {
		description += ", with scattered clouds"
	} else {
		description += ", and cloudy"
	}

	if weather.Rain.OneH > 0 {
		description += ", with recent rain"
	}

	if windSpeed := weather.Wind.Speed; windSpeed > 10 {
		description += ", and windy"
	}

	return description + "."
}
