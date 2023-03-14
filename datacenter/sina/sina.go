package sina

import (
	"encoding/json"
	"fmt"
	goutils "github.com/pocker-lab/goutils"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type JsonSinaStruct struct {
	TotalNum   int        `json:"total_num"`  // 基金总数
	Data       []metadata `json:"data"`       // 基金数据结构体切片
	Lastupdate string     `json:"lastupdate"` // 最后更新时间

	// ExecTime   float64 `json:"exec_time"`  // 执行时间
	// SortTime   float64 `json:"sort_time"`  // 排序时间
}
type metadata struct {
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

// GetSinaData 获取sina网上基金数据；每页请求100条数据
//
// `page`: 页数；
//
// `delay`: 延迟；
func GetSinaData(page, delay int) (data JsonSinaStruct) {
	var (
		tempdata JsonSinaStruct
		client   = &http.Client{Timeout: 30 * time.Second} // 设置http客户端，并设置超时为30秒
		mutex    sync.Mutex                                // 创建一个 mutex，用来保护共享变量 resps
		wg       sync.WaitGroup                            // 创建一个 WaitGroup，用来等待所有的 goroutine 结束
	)

	ch := make(chan struct{}, 5) // 定义一个缓冲大小为10的通道，用于控制并发数量

	// 定义一个字符串变量urls，用来存储请求的URL
	urls := "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['9o_rfPFvmkgcHnSk']/NetValueReturn_Service.NetValueReturnOpen?"

	req, err := http.NewRequest("GET", urls, nil) // 创建一个GET方法的请求，URL为urls，主体为空，并赋值给变量req和err
	goutils.CheckError(err)

	// Params 参数
	q := req.URL.Query() // 调用req.URL.Query()方法，获取请求的查询参数，并赋值给变量q

	// 调用q.Add方法，给查询参数添加多个键值对
	q.Add("page", "1")
	q.Add("num", "100")
	q.Add("sort", "zmjgm")
	q.Add("asc", "0")
	q.Add("ccode", "")
	q.Add("type2", "")
	q.Add("type3", "")

	for i := 165; i <= page; i++ {
		wg.Add(1) // 每次启动一个协程，就增加等待组的计数

		go func(i int) {
			defer wg.Done()  // 协程结束时，减少等待组的计数
			ch <- struct{}{} // 向通道发送一个空结构体，如果通道已满，则阻塞等待

			time.Sleep(time.Duration(delay) * time.Second) // 延迟 5 秒执行
			mutex.Lock()                                   // 加锁
			q.Set("page", strconv.Itoa(i))                 // 调用q.Set方法，设置查询参数page的值为i转换成字符串的结果
			req.URL.RawQuery = q.Encode()                  // 调用q.Encode()方法，把q转换成一个字符串，并赋值给req.URL.RawQuery
			mutex.Unlock()                                 // 解锁

			resp, err := client.Do(req) // 调用client.Do(req)方法，用client发送req这个请求，并返回resp这个响应，并赋值给变量resp和err
			goutils.CheckError(err)

			// 关闭响应体，释放连接资源
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				goutils.CheckError(err)
			}(resp.Body)

			body, err := io.ReadAll(resp.Body) // 调用io.ReadAll(resp.Body)方法，读取响应主体，并赋值给变量body和err
			goutils.CheckError(err)

			err = json.Unmarshal(body[91:len(body)-2], &tempdata) // 将网页内容转换为json格式
			goutils.CheckError(err)

			mutex.Lock()                                    // 加锁
			data.Data = append(data.Data, tempdata.Data...) // 将获取到的内容累加到data结构体
			mutex.Unlock()                                  // 解锁

			fmt.Printf("\r-->  总共有%v条数据，已获取%v条数据", tempdata.TotalNum, len(data.Data))

			<-ch // 从通道接收一个空结构体，释放一个缓冲位置

			data.TotalNum = tempdata.TotalNum     // 基金总数
			data.Lastupdate = tempdata.Lastupdate // 更新时间

			//if tempdata.Data == nil { // 如果获取的内容为空，那么结束循环
			//	//fmt.Println("\n为空")
			//	return

		}(i)
	}

	wg.Wait() // 等待所有协程完成
	fmt.Println()
	return data
}
