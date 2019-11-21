package weather_provider

import (
	"fmt"
	"interface-testing/api/clients/restclient"
	"interface-testing/api/domain/weather_domain"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	restclient.StartMockups()
// 	os.Exit(m.Run())
// }

// var (
// 	getRequestFunc func(url string) (*http.Response, error)
// )

// type getClientMock struct{}

// //We are mocking the service method "Get"
// func (c *getClientMock) Get(request string) (*http.Response, error) {
// 	return getRequestFunc(request)
// }

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type getClientMock struct{}

//We are mocking the service method "Get"
func (cm *getClientMock) Get(request string) (*http.Response, error) {
	return getRequestFunc(request)
}

//When the everything is good
func TestGetWeatherNoError(t *testing.T) {
	// The error we will get is from the "response" so we make the second parameter of the function is nil
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": 44.3601, "longitude": -71.0589, "timezone": "America/New_York", "currently": {"summary": "Clear", "temperature": 40.22, "dewPoint": 50.22, "pressure": 12.90, "humidity": 16.54}}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
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
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 403, "error": "permission denied"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	fmt.Println("this is the error here: ", err)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusForbidden, err.Code)
	assert.EqualValues(t, "permission denied", err.ErrorMessage)
}

func TestGetWeatherInvalidLatitude(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 34223.3445, -71.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "The given location is invalid", err.ErrorMessage)
}

func TestGetWeatherInvalidLongitude(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "The given location is invalid"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -74331.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "The given location is invalid", err.ErrorMessage)
}

func TestGetWeatherInvalidFormat(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "Poorly formatted request"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 0, -74331.0589})

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "Poorly formatted request", err.ErrorMessage)
}

//When no body is provided
func TestGetWeatherInvalidRestClient(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "invalid rest client response"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 0, -74331.0589})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "invalid rest client response", err.ErrorMessage)
}

func TestGetWeatherInvalidResponseBody(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "error": "Invalid response body"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "Invalid response body", err.ErrorMessage)
}

func TestGetWeatherInvalidRequest(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		invalidCloser, _ := os.Open("-asf3")
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       invalidCloser,
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"wrong_anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "invalid argument", err.ErrorMessage)
}

//When the error response is invalid, here the code is supposed to be an integer, but a string was given.
//This can happen when the api owner changes some data types in the api
func TestGetWeatherInvalidErrorInterface(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"code": "string code"}`)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid json response body", err.ErrorMessage)
}

//We are getting a postive response from the api, but, the datatype of the response returned does not match the struct datatype we have defined (does not match the struct type we want to unmarshal this response into).
func TestGetWeatherInvalidResponseInterface(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"latitude": "string latitude", "longitude": -71.0589, "timezone": "America/New_York"}`)), //when we use string for latitude instead of float
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := WeatherProvider.GetWeather(weather_domain.WeatherRequest{"anything", 44.3601, -71.0589})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "error unmarshaling weather fetch response", err.ErrorMessage)
}
