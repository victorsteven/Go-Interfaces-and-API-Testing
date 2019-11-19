package weather_domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeather(t *testing.T) {
	request := Weather{
		Latitude:  12.33,
		Longitude: 90.34,
		TimeZone:  "America/New_York",
		Currently: CurrentlyInfo{
			Temperature: 10,
			Summary:     "Clear",
			DewPoint:    20.433,
			Pressure:    95.33,
			Humidity:    71.34,
		},
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var result Weather
	err = json.Unmarshal(bytes, &result)

	assert.Nil(t, err)
	assert.EqualValues(t, result.Latitude, request.Latitude)
	assert.EqualValues(t, result.TimeZone, request.TimeZone)
	assert.EqualValues(t, result.Longitude, request.Longitude)
	assert.EqualValues(t, result.Currently.Summary, request.Currently.Summary)
	assert.EqualValues(t, result.Currently.Humidity, request.Currently.Humidity)
	assert.EqualValues(t, result.Currently.DewPoint, request.Currently.DewPoint)
	assert.EqualValues(t, result.Currently.Pressure, request.Currently.Pressure)
	assert.EqualValues(t, result.Currently.Temperature, request.Currently.Temperature)
}

func TestWeatherError(t *testing.T) {
	request := WeatherError{
		Code:         400,
		ErrorMessage: "Bad Request Error",
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var errResult WeatherError
	err = json.Unmarshal(bytes, &errResult)
	assert.Nil(t, err)
	assert.EqualValues(t, errResult.Code, request.Code)
	assert.EqualValues(t, errResult.ErrorMessage, request.ErrorMessage)

}