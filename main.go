package main

import (
	"log"
	"spider"
)

func main() {
	//citySpider := spider.NewCitySpider()
	//citySpider.GetCityInfo("彭州")
	weatherSpider := spider.NewWeatherSpider()
	weatherSpider.GetWeatherInfo("101010100")
	log.Println(*weatherSpider.PartsWeather)
}
