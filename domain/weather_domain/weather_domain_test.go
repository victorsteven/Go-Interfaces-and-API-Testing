package weather_domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeather(t *testing.T) {
	request := Weather{
		Latitude:     12.33,
		Longitude:    90.34,
		TimeZone:     "America/New_York",
		CurrentState: Currently{
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
	assert.EqualValues(t,result.Latitude, request.Latitude)
	assert.EqualValues(t,result.TimeZone, request.TimeZone)
	assert.EqualValues(t,result.Longitude, request.Longitude)
	assert.EqualValues(t,result.CurrentState.Summary, request.CurrentState.Summary)
	assert.EqualValues(t,result.CurrentState.Humidity, request.CurrentState.Humidity)
	assert.EqualValues(t,result.CurrentState.DewPoint, request.CurrentState.DewPoint)
	assert.EqualValues(t,result.CurrentState.Pressure, request.CurrentState.Pressure)
	assert.EqualValues(t,result.CurrentState.Temperature, request.CurrentState.Temperature)
}