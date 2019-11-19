package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func RunApp(){

	routes()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

