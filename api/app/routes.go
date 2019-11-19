package app

import "interface-testing/api/controllers/weather_controller"

func routes() {
	router.GET("/weather/:apiKey/:latitude/:longitude", weather_controller.GetWeather)
}
