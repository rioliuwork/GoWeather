package services

import (
	"GoSpider/spider"
	"github.com/gin-gonic/gin"
)

func CityList(c *gin.Context) {
	cityName := c.PostForm("cityName")
	citySpider := spider.NewCitySpider()
	citySpider.GetCityInfo(cityName)
	c.SecureJSON(200, gin.H{
		"status": "success",
		"citys":  citySpider.Citys,
	})
}
