package weather_provider

import (
	"encoding/json"
	"fmt"
	"interface-testing/clients"
	"interface-testing/domain/weather_domain"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	weatherUrl = "https://api.darksky.net/forecast/%s/%f,%f"
)

func GetWeatherInfo(accessToken string, latitude float64, longitude float64) (*weather_domain.Weather, *weather_domain.WeatherError){
	url := fmt.Sprintf(weatherUrl, accessToken, latitude, longitude)
	response, err := clients.GetInfo(url)
	if err != nil {
		return nil, &weather_domain.WeatherError{
			Code: http.StatusInternalServerError,
			Error: "THis is the error",
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &weather_domain.WeatherError{
			Code:  http.StatusInternalServerError,
			Error: "THis is the second error",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse weather_domain.WeatherError
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &weather_domain.WeatherError{
				Code:  http.StatusInternalServerError,
				Error: "Invalid json response body",
			}
		}
		errResponse.Code = response.StatusCode
		return nil, &errResponse
	}
	var result weather_domain.Weather
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal dark sky response: %s", err.Error()))
		return nil, &weather_domain.WeatherError{Code: http.StatusInternalServerError, Error: "error when trying to unmarshal darksky fetch data response"}
	}
	return &result, nil
}

