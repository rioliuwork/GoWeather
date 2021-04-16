package services

import (
	"GoSpider/spider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DayAndNight(c *gin.Context) {
	htmlCode := c.PostForm("htmlCode")
	weatherSpider := spider.NewWeatherSpider()
	weatherSpider.GetWeatherInfo(htmlCode)
	c.SecureJSON(http.StatusOK, gin.H{
		"status":       "success",
		"dayWeather":   (*weatherSpider.PartsWeather)[0],
		"nightWeather": (*weatherSpider.PartsWeather)[1],
	})
}
