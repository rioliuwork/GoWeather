package spider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type City struct {
	HtmlCode     string
	ProvincePY   string
	CityName     string
	CityPY       string
	CityCode     string
	ProvinceName string
}

type CitySpider struct {
	Citys []City
}

func (spider *CitySpider) GetCityInfo(cityName string) {
	jsonArray := search(cityName)
	citys := make([]City, 0)
	for _, cityInfoMap := range jsonArray {
		infos := strings.Split(cityInfoMap["ref"].(string), "~")
		city := City{
			HtmlCode:     infos[0],
			ProvincePY:   infos[1],
			CityName:     infos[2],
			CityPY:       infos[3],
			CityCode:     infos[8],
			ProvinceName: infos[9],
		}
		citys = append(citys, city)
	}
	spider.Citys = citys
}

//查询城市列表
func search(cityName string) (jsonArray []map[string]interface{}) {
	reqUrl := "http://toy1.weather.com.cn/search"
	param := "cityname=" + cityName
	cookie := &http.Cookie{Name: "csrfToken", Value: "ou8D3UP4LRNKgk5iiEEpRU2s", Domain: "toy1.weather.com.cn", Path: "/",
		Expires: time.Now().Add(111 * time.Second), Secure: false, HttpOnly: false}
	res, _ := GetNetResponse(reqUrl, &param, cookie)
	defer res.Body.Close()
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyByte)
	bodyJson := body[1 : len(body)-1]
	json.Unmarshal([]byte(bodyJson), &jsonArray)
	log.Println(jsonArray)
	return
}
