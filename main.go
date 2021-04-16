package main

import (
	"GoSpider/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/city/list", func(context *gin.Context) {
		services.CityList(context)
	})
	router.POST("/weather/difference", func(context *gin.Context) {
		services.DayAndNight(context)
	})
	router.Run(":8080")
}
