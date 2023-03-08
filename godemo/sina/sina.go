package sina

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetSina(num int) (str string) {
	url := "https://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList%5B%25279o_rfPFvmkgcHnSk%2527%5D/NetValueReturn_Service.NetValueReturnOpen?"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "vip.stock.finance.sina.com.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "MONEY-FINANCE-SINA-COM-CN-WEB5=")
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(num))
	q.Add("num", "9000")
	q.Add("sort", "zmjqm")
	q.Add("asc", "0")
	q.Add("ccode", "")
	q.Add("type2", "")
	q.Add("type3", "")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	str = string(body[91 : len(body)-2])
	return str
}
