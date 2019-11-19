package weather_domain

import (
	"encoding/json"
	"net/http"
)

type WeatherErrorInterface interface {
	Status() int
	Message() string
}
type WeatherError struct {
	Code      int           `json:"code"`
	ErrorMessage     string        `json:"error"`
}

func (w *WeatherError) Status() int {
	return w.Code
}
func (w *WeatherError) Message() string {
	return w.ErrorMessage
}

func NewWeatherError(statusCode int, message string) WeatherErrorInterface {
	return &WeatherError{
		Code:         statusCode,
		ErrorMessage: message,
	}
}
func NewBadRequestError(message string) WeatherErrorInterface {
	return &WeatherError{
		Code: http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewForbiddenError(message string) WeatherErrorInterface {
	return &WeatherError{
		Code: http.StatusForbidden,
		ErrorMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (WeatherErrorInterface, error) {
	var result WeatherError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}


