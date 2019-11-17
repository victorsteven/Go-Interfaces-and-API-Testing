package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"interface-testing/api/domain/weather_domain"
	"interface-testing/api/providers/weather_provider"
	"log"
)

func init(){
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
func main(){
 	//response, err :=  weather_provider.GetWeather(os.Getenv("DARK_SKY_SECRET_API_KEY"), 42.3601, -71.0589)
	response, err :=  weather_provider.GetWeather(weather_domain.WeatherRequest{"4234wdsfsdf34234", 42.3601, -71.0589})
	if err != nil {
 		fmt.Println("This is the error: ", err)
 		return
	}
	fmt.Println("This is the response latitude: ", response.Latitude)
	fmt.Println("This is the response longitude: ", response.Longitude)
	//fmt.Println("This is the response currently: ", response.Currently)
}
