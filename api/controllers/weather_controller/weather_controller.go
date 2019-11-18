package weather_controller

import (
	"github.com/gin-gonic/gin"
	"interface-testing/api/domain/weather_domain"
	"interface-testing/api/services"
	"net/http"
	"strconv"
)
func GetWeather(c *gin.Context){

	lat, err := strconv.ParseFloat(c.Param("latitude"), 64)
	if err != nil {
		apiError := weather_domain.WeatherErrorResponse{Code: http.StatusBadRequest, ErrorMessage:"invalid latitude body"}
		c.JSON(apiError.Code, apiError)
		return
	}
	long, err := strconv.ParseFloat(c.Param("longitude"), 64)
	if err != nil {
		apiError := weather_domain.WeatherErrorResponse{Code: http.StatusBadRequest, ErrorMessage:"invalid longitude body"}
		c.JSON(apiError.Code, apiError)
		return
	}
	request :=  weather_domain.WeatherRequest{
		ApiKey:    c.Param("apiKey"),
		Latitude:  lat,
		Longitude: long,
	}
	//if err := c.ShouldBindJSON(&request); err != nil {
	//	//	apiError := weather_domain.WeatherErrorResponse{ErrorMessage:"invalid json body"}
	//	//	c.JSON(apiError.Code, apiError)
	//	//	return
	//	//}
	result, apiError := services.WeatherService.GetWeather(request)
	if apiError != nil {
		c.JSON(apiError.Status(), apiError)
		return
	}
	c.JSON(http.StatusOK, result)
}