package services

import (
	"github.com/stretchr/testify/assert"
	"interface-testing/api/clients/restclient"
	"interface-testing/api/domain/weather_domain"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestWeatherServiceNoAuthKey(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/wrong_key/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusForbidden,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 403, "error": "permission denied"}`)),
		},
	})
	request := weather_domain.WeatherRequest{ApiKey: "wrong_key", Latitude: 44.3601, Longitude: -71.0589}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t,result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusForbidden, err.Status())
	assert.EqualValues(t, "permission denied", err.Message())
}

func TestWeatherServiceWrongLatitude(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/api_key/123443,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		},
	})
	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 123443, Longitude: -71.0589}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t,result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "The given location is invalid", err.Message())
}

func TestWeatherServiceWrongLongitude(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/api_key/39.12,122332",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		},
	})
	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 39.12, Longitude: 122332}
	result, err := WeatherService.GetWeather(request)
	assert.Nil(t,result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "The given location is invalid", err.Message())
}

func TestWeatherServiceSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/api_key/39.12,49.12",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": 39.12, "longitude": 49.12, "timezone": "America/New_York", "currently": {"summary": "Clear", "temperature": 40.22, "dewPoint": 50.22, "pressure": 12.90, "humidity": 16.54}}`)),
		},
	})
	request := weather_domain.WeatherRequest{ApiKey: "api_key", Latitude: 39.12, Longitude: 49.12}
	result, err := WeatherService.GetWeather(request)
	assert.NotNil(t,result)
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