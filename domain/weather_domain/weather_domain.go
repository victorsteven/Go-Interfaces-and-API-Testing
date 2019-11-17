package weather_domain

type Weather struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"timezone"`
	Currently CurrentlyInfo `json:"currently"`
}

type CurrentlyInfo struct {
	Temperature float64 `json:"temperature"`
	Summary string `json:"summary"`
	DewPoint float64 `json:"dewPoint"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

type WeatherRequest struct {
	ApiKey string `json:"api_key"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}


