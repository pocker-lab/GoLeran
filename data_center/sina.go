package data_center

import (
	"fmt"
	"net/http"
	"strconv"
)

type Sina struct {
	symbol string // 基金代码
	sname  string // 基金名称
}

func SinaGetRequest(num int) (sinaData string) {
	req, err := http.NewRequest("GET", "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList%5B%2527uUkyqh1ha4F3vGIy%2527%5D/NetValueReturn_Service.NetValueReturnOpen?", nil)
	if err != nil {
		fmt.Println(err)
	}
	numstr := strconv.Itoa(num)
	//q.Add("page", numstr)
	//q.Add("num", "2")
	//q.Add("sort", "zmjgm")
	//q.Add("asc", "0")

	req.Header.Add("page", numstr)

	fmt.Println(req.Body)
	fmt.Println()
	fmt.Println(numstr)
	return sinaData
}
