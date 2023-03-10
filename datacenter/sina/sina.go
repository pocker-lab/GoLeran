package sina

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	goutils "github.com/pocker-lab/goutils"
)

type sinaJson struct {
	TotalNum   int    `json:"total_num"`  // 基金总数
	Data       []data `json:"data"`       // 基金数据结构体切片
	Lastupdate string `json:"lastupdate"` // 最后更新时间

	// ExecTime   float64 `json:"exec_time"`  // 执行时间
	// SortTime   float64 `json:"sort_time"`  // 排序时间
}
type data struct {
	Symbol interface{} `json:"symbol"` // 基金代码
	Name   string      `json:"name"`   // 基金全称
	Clrq   string      `json:"clrq"`   // 成立日期
	Jjjl   string      `json:"jjjl"`   // 基金经理

	//Sname     string      `json:"sname"`      // 基金简称
	//Zmjgm     string      `json:"zmjgm"`      // 总募集规模(万份)
	//JjglrCode string      `json:"jjglr_code"` // 基金管理人代码
	// PerNav     string      `json:"per_nav"`     // 单位净值(元)
	// TotalNav   string      `json:"total_nav"`   // 七日年化收益率(%)
	// ThreeMonth float64     `json:"three_month"` // 近3月涨幅
	// SixMonth   float64     `json:"six_month"`   // 近6月涨幅
	// OneYear    float64     `json:"one_year"`    // 近1年涨幅
	// FormYear   float64     `json:"form_year"`
	// FormStart  float64     `json:"form_start"`
	// Dwjz       string      `json:"dwjz"`       // 万份收益
	// Ljjz       string      `json:"ljjz"`       // 7日年化
	// Jzrq       string      `json:"jzrq"`       // 截止日期
	// Zjzfe      int         `json:"zjzfe"`      // 最近总份额(万份)
}

func MainSina() {
	str1 := GetSina(1, 9000)
	str2 := GetSina(2, 9000)
	var (
		au1 = sinaJson{}
		au2 = sinaJson{}
	)

	err := json.Unmarshal([]byte(str1), &au1)
	goutils.CheckError(err)
	err = json.Unmarshal([]byte(str2), &au2)
	goutils.CheckError(err)

	au1.Data = append(au1.Data, au2.Data...)
	fmt.Printf("%v--->%v\n", au1.TotalNum, len(au1.Data))

	// 将数据写入到文件中
	bytes, _ := json.MarshalIndent(au1.Data, "", "  ")
	os.WriteFile("sina.json", bytes, 0644)
	//fmt.Printf(string(bytes))
	//Goutils.WriteFile2(string(bytes))

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/godb?charset=utf8")
	goutils.CheckError(err)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(db)

	_, err = db.Exec("ALTER TABLE sina AUTO_INCREMENT = 0")
	goutils.CheckError(err)

	tx, err := db.Begin()
	goutils.CheckError(err)

	//result, err := tx.Exec("INSERT INTO godb.sina (symbol,name,clrq,jjjl) VALUES(?,?,?,?)", "000001", "华夏成长混合", "2001-12-18 00:00:00", "王泽实、万方方")
	//goutils.CheckError(err)
	//fmt.Println(result.RowsAffected())

	for _, k := range au1.Data {
		stmt, err := tx.Prepare("INSERT INTO godb.sina (symbol,name,clrq,jjjl) VALUES(?,?,?,?)")
		goutils.CheckError(err)

		_, err = stmt.Exec(k.Symbol, k.Name, k.Clrq, k.Jjjl)
		goutils.CheckError(err)

		err = stmt.Close()
		goutils.CheckError(err)
	}
	defer tx.Commit()
}

func GetSina(page, num int) (str string) {
	url := "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['9o_rfPFvmkgcHnSk']/NetValueReturn_Service.NetValueReturnOpen"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	goutils.CheckError(err)

	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "vip.stock.finance.sina.com.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "MONEY-FINANCE-SINA-COM-CN-WEB5=")
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("num", strconv.Itoa(num))
	q.Add("sort", "zmjqm")
	q.Add("asc", "0")
	q.Add("ccode", "")
	q.Add("type2", "")
	q.Add("type3", "")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	goutils.CheckError(err)
	body, err := io.ReadAll(res.Body)
	goutils.CheckError(err)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	//str = string(body[91 : len(body)-2])
	str = string(body[91 : len(body)-2])
	return str
}
