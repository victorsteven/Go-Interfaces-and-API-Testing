package weather_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"interface-testing/clients/restclient"
	"interface-testing/domain/weather_domain"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

//When the everything is good
func TestGetWeatherNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/44.3601,-71.0589", //this should match with the parameters supplied to the GetWeather below
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": 44.3601, "longitude": -71.0589, "timezone": "America/New_York", "currently": {"summary": "Clear", "temperature": 40.22, "dewPoint": 50.22, "pressure": 12.90, "humidity": 16.54}}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 44.3601, response.Latitude)
	assert.EqualValues(t, -71.0589, response.Longitude)
	assert.EqualValues(t, "America/New_York", response.TimeZone)
	assert.EqualValues(t, "Clear", response.Currently.Summary)
	assert.EqualValues(t, 40.22, response.Currently.Temperature)
	assert.EqualValues(t, 50.22, response.Currently.DewPoint)
	assert.EqualValues(t, 16.54, response.Currently.Humidity)
	assert.EqualValues(t, 12.90, response.Currently.Pressure)
}

func TestGetWeatherInvalidApiKey(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/wrong_anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusForbidden,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 403, "error": "permission denied"}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusForbidden, err.Code)
	assert.EqualValues(t, "permission denied", err.Error)
}


func TestGetWeatherInvalidLatitude(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/34223.3445,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 34223.3445, -71.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "The given location is invalid", err.Error)
}

func TestGetWeatherInvalidLongitude(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/44.3601,-74331.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -74331.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "The given location is invalid", err.Error)
}

func TestGetWeatherInvalidFormat(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/0,-74331.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "Poorly formatted request"}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 0, -74331.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "Poorly formatted request", err.Error)
}


//When no body is provided
func TestGetWeatherInvalidRestClient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/0,-74331.0589",
		HttpMethod: http.MethodGet,
		Err: errors.New("invalid rest client response"),
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 0, -74331.0589})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "invalid rest client response", err.Error)
}

func TestGetWeatherInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/wrong_anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Err: errors.New("Invalid response body"),
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "Invalid response body", err.Error)
}

func TestGetWeatherInvalidRequest(t *testing.T) {
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/wrong_anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       invalidCloser,
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "invalid argument", err.Error)
}

//When the error response is invalid, here the code is supposed to be an integer, but a string was given.
//This can happen when the api owner changes some data types in the api
func TestGetWeatherInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:  ioutil.NopCloser(strings.NewReader(`{"code": "string code"}`)),
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid json response body", err.Error)
}

//We are getting a postive response from the api, but, the datatype of the response returned does not match the struct datatype we have defined (does not match the struct type we want to unmarshal this response into).
func TestGetWeatherInvalidResponseInterface(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": "string latitude", "longitude": -71.0589, "timezone": "America/New_York"}`)), //when we use string for latitude instead of float
		},
	})
	response, err := GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "error unmarshaling weather fetch response", err.Error)
}


