package spider

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetNetResponse(url string, param *string, cookie *http.Cookie) (*http.Response, error) {
	client := &http.Client{}
	var req *http.Request
	var searchUrl string
	if param != nil {
		searchUrl = strings.Join([]string{url, getParseParam(param)}, "?")
	} else {
		searchUrl = url
	}
	req, _ = http.NewRequest("GET", searchUrl, nil)
	if cookie != nil {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return res, nil
}

//将get请求的参数进行转义
func getParseParam(param *string) string {
	return url.PathEscape(*param)
}
