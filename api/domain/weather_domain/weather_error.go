package weather_domain

import (
	"encoding/json"
	"net/http"
)

type WeatherErrorResponse struct {
	Code      int           `json:"code"`
	ErrorMessage     string        `json:"error"` //this is the response json
}
type weatherError struct {
	Code      int           `json:"code"`
	ErrorMessage     string        `json:"error"` //this is the response json
}
type WeatherErrorInterface interface {
	Status() int
	Message() string
}
func (w *weatherError) Status() int {
	return w.Code
}
func (w *weatherError) Message() string {
	return w.ErrorMessage
}

func NewWeatherError(statusCode int, message string) WeatherErrorInterface {
	return &weatherError{
		Code:         statusCode,
		ErrorMessage: message,
	}
}
func NewBadRequestError(message string) WeatherErrorInterface {
	return &weatherError{
		Code: http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewForbiddenError(message string) WeatherErrorInterface {
	return &weatherError{
		Code: http.StatusForbidden,
		ErrorMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (WeatherErrorInterface, error) {
	var result weatherError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}


