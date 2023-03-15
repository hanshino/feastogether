package fetch

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func Post(api string, data []byte, act string, token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Failed to create http request object: ", err)
		return nil, err
	}

	// add header
	addHeader(req, act, token)

	if resp, err := client.Do(req); err != nil {
		log.Println("Failed to execute http request: ", err)
		return nil, err
	} else {
		return resp, err
	}
}

func addHeader(req *http.Request, act string, token string) *http.Request {

	if token != "" {
		req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Add("authority", "www.feastogether.com.tw")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("act", act)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://www.feastogether.com.tw")
	req.Header.Add("referer", "https://www.feastogether.com.tw/booking/Eatogether/search")
	req.Header.Add("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Windows"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	return req
}
