package data_center

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// http://vip.stock.finance.sina.com.cn/fund_center/index.html#jjgmall

// http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['uUkyqh1ha4F3vGIy']/NetValueReturn_Service.NetValueReturnOpen?

func SinaGetRequest(num int) (sinaData string) {
	req, err := http.NewRequest("GET", "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList%5B%2527uUkyqh1ha4F3vGIy%2527%5D/NetValueReturn_Service.NetValueReturnOpen?", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	numstr := strconv.Itoa(num)
	q.Add("page", numstr)
	q.Add("num", "2")
	q.Add("sort", "zmjgm")
	q.Add("asc", "0")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	sinaData = string(body)
	return sinaData
}
