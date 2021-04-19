package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Weather struct {
	PartsWeather *[]partWeather
}

type partWeather struct {
	//时间名称
	TimeName string
	//天气
	Wea string
	//温度
	Tem string
	//风向
	Win string
	//日出时间
	SunUp string
	//日落时间
	SunDown string
}

type LifeInfo struct {
	IndexName string
	Title     string
	Context   string
}

type WeatherSpider struct {
	PartsWeather *[]partWeather
	LifeInfos    *[]LifeInfo
}

func (spider *WeatherSpider) GetWeatherInfo(htmlCode string) {
	reqUrl := fmt.Sprintf("http://www.weather.com.cn/weather1d/%s.shtml", htmlCode)
	res, _ := GetNetResponse(reqUrl, nil, nil)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	spider.PartsWeather = parsePartsWeather(doc)
	spider.LifeInfos = parseLifeInfos(doc)
	//log.Println(partsWeather)

	//fmt.Printf("script:%s",doc.Find("script").Text())
}

func parseLifeInfos(doc *goquery.Document) *[]LifeInfo {
	lifeInfos := make([]LifeInfo, 0)
	doc.Find(".livezs .clearfix li").Each(func(i int, selection *goquery.Selection) {
		lifeInfo := LifeInfo{
			Title:     strings.Replace(selection.Find("span").Text(), "\n", "", -1),
			IndexName: selection.Find("em").Text(),
			Context:   selection.Find("p").Text(),
		}
		lifeInfos = append(lifeInfos, lifeInfo)
	})
	return &lifeInfos
}

func parsePartsWeather(doc *goquery.Document) *[]partWeather {
	partsWeather := make([]partWeather, 0)
	doc.Find(".t .clearfix li:not(.lv1):not(.lv2):not(.lv3):not(.lv4)").Each(func(i int, selection *goquery.Selection) {
		winVal, _ := selection.Find(".win span").Attr("title")
		win := strings.Join([]string{winVal, selection.Find(".win span").Text()}, " ")
		tem := fmt.Sprintf("%s%s", selection.Find(".tem span").Text(), selection.Find(".tem em").Text())
		sunUp := selection.Find(".sunUp span").Text()
		sunDowm := selection.Find(".sunDown span").Text()
		partW := partWeather{
			TimeName: selection.Find("h1").Text(),
			Wea:      selection.Find(".wea").Text(),
			Tem:      tem,
			Win:      win,
			SunUp:    strings.Replace(sunUp, "\n", " ", -1),
			SunDown:  strings.Replace(sunDowm, "\n", " ", -1),
		}
		partsWeather = append(partsWeather, partW)
	})
	return &partsWeather
}
