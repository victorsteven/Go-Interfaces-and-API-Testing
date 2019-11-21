package services

import (
	"interface-testing/api/domain/weather_domain"
	"interface-testing/api/providers/weather_provider"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getWeatherProviderFunc func(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError)
)
type getProviderMock struct{}

func (c *getProviderMock) GetWeather(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError) {
	return getWeatherProviderFunc(request)
}

func TestWeatherServiceNoAuthKey(t *testing.T) {
	getWeatherProviderFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError) {
		return nil, &weather_domain.WeatherError{
			Code:         403,
			ErrorMessage: "permission denied",
		}
	}
	weather_provider.WeatherProvider = &getProviderMock{} //without this line, the real api is fired

	request := weather_domain.WeatherRequest{ApiKey: "wrong_key", Latitude: 44.3601, Longitude: -71.0589}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusForbidden, err.Status())
	assert.EqualValues(t, "permission denied", err.Message())
}

func TestWeatherServiceWrongLatitude(t *testing.T) {
	getWeatherProviderFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError) {
		return nil, &weather_domain.WeatherError{
			Code:         400,
			ErrorMessage: "The given location is invalid",
		}
	}
	weather_provider.WeatherProvider = &getProviderMock{} //without this line, the real api is fired

	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 123443, Longitude: -71.0589}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "The given location is invalid", err.Message())
}

func TestWeatherServiceWrongLongitude(t *testing.T) {
	getWeatherProviderFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError) {
		return nil, &weather_domain.WeatherError{
			Code:         400,
			ErrorMessage: "The given location is invalid",
		}
	}
	weather_provider.WeatherProvider = &getProviderMock{} //without this line, the real api is fired

	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 39.12, Longitude: 122332}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "The given location is invalid", err.Message())
}

func TestWeatherServiceSuccess(t *testing.T) {
	getWeatherProviderFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather, *weather_domain.WeatherError) {
		return &weather_domain.Weather{
			Latitude:  39.12,
			Longitude: 49.12,
			TimeZone:  "America/New_York",
			Currently: weather_domain.CurrentlyInfo{
				Temperature: 40.22,
				Summary:     "Clear",
				DewPoint:    50.22,
				Pressure:    12.90,
				Humidity:    16.54,
			},
		}, nil
	}
	weather_provider.WeatherProvider = &getProviderMock{} //without this line, the real api is fired

	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 39.12, Longitude: 49.12}
	result, err := WeatherService.GetWeather(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 39.12, result.Latitude)
	assert.EqualValues(t, 49.12, result.Longitude)
	assert.EqualValues(t, "America/New_York", result.TimeZone)
	assert.EqualValues(t, "Clear", result.Currently.Summary)
	assert.EqualValues(t, 40.22, result.Currently.Temperature)
	assert.EqualValues(t, 50.22, result.Currently.DewPoint)
	assert.EqualValues(t, 12.90, result.Currently.Pressure)
	assert.EqualValues(t, 16.54, result.Currently.Humidity)
}
