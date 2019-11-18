package weather_domain

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


