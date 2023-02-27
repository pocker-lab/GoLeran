package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "http://api.fund.eastmoney.com/f10/lsjz?callback=jQuery18308421360004753573_1677422025866&fundCode=320007&pageIndex=2&pageSize=20"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "api.fund.eastmoney.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Referer", "http://fundf10.eastmoney.com/jjjz_320007.html")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
