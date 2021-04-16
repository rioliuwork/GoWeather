package main

import (
	"log"
	"spider"
)

func main() {
	citySpider := spider.NewCitySpider()
	citySpider.GetCityInfo("四川")
	weatherSpider := spider.NewWeatherSpider()
	weatherSpider.GetWeatherInfo(citySpider.Citys[0].HtmlCode)
	log.Println(*weatherSpider.PartsWeather)
}
