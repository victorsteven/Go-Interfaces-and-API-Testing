package weather_provider

import (
	"github.com/stretchr/testify/assert"
	"interface-testing/clients/restclient"
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

func TestGetWeatherNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.darksky.net/forecast/anything/44.3601,-71.0589",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": 44.3601, "longitude": -71.0589, "timezone": "America/New_York", "currently": {"summary": "Clear", "temperature": 40.22, "dewPoint": 50.22, "pressure": 12.90, "humidity": 16.54}}`)),
		},
	})
	response, err := GetWeather("anything", 44.3601, -71.0589)
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
	response, err := GetWeather("wrong_anything", 44.3601, -71.0589)

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
	response, err := GetWeather("anything", 34223.3445, -71.0589)

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
	response, err := GetWeather("anything", 44.3601, -74331.0589)

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
	response, err := GetWeather("anything", 0, -74331.0589)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "Poorly formatted request", err.Error)
}

