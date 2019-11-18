package weather_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interface-testing/api/clients/restclient"
	"interface-testing/api/domain/weather_domain"
	"interface-testing/api/services"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}

var (
	getWeatheFunc func(request weather_domain.WeatherRequest) (*weather_domain.Weather, weather_domain.WeatherErrorInterface)
)

type weatherServiceMock struct {}

//We are mocking the service method "GetWeather"
func (w *weatherServiceMock) GetWeather(request weather_domain.WeatherRequest) (*weather_domain.Weather, weather_domain.WeatherErrorInterface) {
	return getWeatheFunc(request)
}


func TestGetWeatherLatitudeInvalid(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return nil, weather_domain.NewBadRequestError("invalid latitude body")
	}
	services.WeatherService = &weatherServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "right_api_key"},
		{Key: "latitude", Value:  "1rte4.78"},
		{Key: "longitude", Value: fmt.Sprintf("%f", 42.78)},
	}
	GetWeather(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := weather_domain.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid latitude body", apiErr.Message())
}

func TestGetWeatherLongitudeInvalid(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return nil, weather_domain.NewBadRequestError("invalid longitude body")
	}
	services.WeatherService = &weatherServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "right_api_key"},
		{Key: "latitude", Value: fmt.Sprintf("%f", 12.78)},
		{Key: "longitude", Value:  "23awe.78"},
	}
	GetWeather(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := weather_domain.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid longitude body", apiErr.Message())
}

func TestGetWeatherLatitudeInvalidLocation(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return nil, weather_domain.NewBadRequestError("The given location is invalid")
	}
	services.WeatherService = &weatherServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "right_api_key"},
		{Key: "latitude", Value: fmt.Sprintf("%f", 122334.78)},
		{Key: "longitude", Value: fmt.Sprintf("%f", 42.78)},
	}
	GetWeather(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := weather_domain.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "The given location is invalid", apiErr.Message())
}

func TestGetWeatherLongitudeInvalidLocation(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return nil, weather_domain.NewBadRequestError("The given location is invalid")
	}
	services.WeatherService = &weatherServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "right_api_key"},
		{Key: "latitude", Value: fmt.Sprintf("%f", 12.78)},
		{Key: "longitude", Value: fmt.Sprintf("%f", 423243.78)},
	}
	GetWeather(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := weather_domain.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "The given location is invalid", apiErr.Message())
}

////When the run status code is supplied
func TestGetWeatherInvalidKey(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return nil, weather_domain.NewForbiddenError("permission denied")
	}
	services.WeatherService = &weatherServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "wrong_api_key"},
		{Key: "latitude", Value: fmt.Sprintf("%f", 12.78)},
		{Key: "longitude", Value: fmt.Sprintf("%f", 42.78)},
	}
	GetWeather(c)
	var apiError weather_domain.WeatherErrorResponse
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusForbidden, response.Code)
	assert.EqualValues(t, "permission denied", apiError.ErrorMessage)
}
//
func TestGetWeatherSuccess(t *testing.T) {
	getWeatheFunc = func(request weather_domain.WeatherRequest) (*weather_domain.Weather,  weather_domain.WeatherErrorInterface) {
		return &weather_domain.Weather{
			Latitude:  20.34,
			Longitude: -12.44,
			TimeZone:  "Africa/Nouakchott",
			Currently: weather_domain.CurrentlyInfo{
				Temperature: 78.02,
				Summary:     "Overcast",
				DewPoint:    32.37,
				Pressure:    1014.1,
				Humidity:    0.19,
			},
		}, nil
	}
	services.WeatherService = &weatherServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "apiKey", Value: "right_api_key"},
		{Key: "latitude", Value: fmt.Sprintf("%f", 20.34)},
		{Key: "longitude", Value: fmt.Sprintf("%f", -12.44)},
	}
	GetWeather(c)
	var weather weather_domain.Weather
	err := json.Unmarshal(response.Body.Bytes(), &weather)
	assert.Nil(t, err)
	assert.NotNil(t, weather)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, 20.34, weather.Latitude)
	assert.EqualValues(t, -12.44, weather.Longitude)
	assert.EqualValues(t, "Africa/Nouakchott", weather.TimeZone)
	assert.EqualValues(t, 78.02, weather.Currently.Temperature)
	assert.EqualValues(t, 0.19, weather.Currently.Humidity)
	assert.EqualValues(t, 1014.1, weather.Currently.Pressure)
	assert.EqualValues(t, 32.37, weather.Currently.DewPoint)
	assert.EqualValues(t, "Overcast", weather.Currently.Summary)
}
