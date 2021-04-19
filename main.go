package main

import (
	"GoSpider/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/city/list", services.CityList)
	router.POST("/weather/difference", services.DayAndNight)
	router.Run(":8080")
}
