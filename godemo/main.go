package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godemo/sina"
	"log"
)

type jsonA struct {
	TotalNum   int    `json:"total_num"`  // 基金总数
	Data       []data `json:"data"`       // 基金数据切片
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

func main() {
	str1 := sina.GetSina(1)
	str2 := sina.GetSina(2)
	var (
		au1 = jsonA{}
		au2 = jsonA{}
	)

	json.Unmarshal([]byte(str1), &au1)
	json.Unmarshal([]byte(str2), &au2)

	au1.Data = append(au1.Data, au2.Data...)
	fmt.Printf("%v--->%v\n", au1.TotalNum, len(au1.Data))

	//jsonStudent, err := json.Marshal(au1)
	//if err != nil {
	//	fmt.Println("转换为json错误")
	//}
	//str3 := jsonStudent
	//Goutils.WriteFile2(string(str3))

	// 将数据写入到文件中
	//bytes, _ := json.MarshalIndent(au1.Data, "", "  ")
	//ioutil.WriteFile("config.json", bytes, 0644)
	//fmt.Printf(string(bytes))
	//Goutils.WriteFile2(string(bytes))

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/godb?charset=utf8")
	check(err)
	defer db.Close()
	db.Exec("ALTER TABLE sina AUTO_INCREMENT = 0")

	tx, err := db.Begin()
	check(err)

	//result, err := tx.Exec("INSERT INTO godb.sina (symbol,name,clrq,jjjl) VALUES(?,?,?,?)", "000001", "华夏成长混合", "2001-12-18 00:00:00", "王泽实、万方方")
	//check(err)
	//fmt.Println(result.RowsAffected())

	for _, k := range au1.Data {
		stmt, err := tx.Prepare("INSERT INTO godb.sina (symbol,name,clrq,jjjl) VALUES(?,?,?,?)")
		check(err)
		stmt.Exec(k.Symbol, k.Name, k.Clrq, k.Jjjl)
		//fmt.Println(res.RowsAffected())
		stmt.Close()
	}

	err = tx.Commit()
	check(err)

}

// check 因为要多次检查错误，所以干脆自己建立一个函数。
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
