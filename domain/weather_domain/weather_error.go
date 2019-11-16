package weather_domain


type WeatherError struct {
	Code      int           `json:"code"`
	Error     string        `json:"error"`
}