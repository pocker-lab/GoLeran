// Package datacenter 数据来源
package datacenter

import (
	"GoLeran/datacenter/sina"
	"encoding/json"
	"fmt"
	goutils "github.com/pocker-lab/goutils"
	"os"
)

func DataCenter() {
	// go goutils.Timing(time.Now())
	// data := sina.GetSinaData(167, 1)
	// fmt.Printf("%+v-->%+v-->%+v", data.TotalNum, data.Lastupdate, len(data.Data))
	// WriteFile(data, "sina")
}

func WriteFile(content sina.JsonSinaStruct, name string) {
	path := fmt.Sprintf("datacenter/%v/%v.json", name, name)
	openFile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE, 0666)
	goutils.CheckError(err)
	defer func(openFile *os.File) {
		err := openFile.Close()
		goutils.CheckError(err)
	}(openFile)

	// 通过 JSON 序列化结构体
	text, _ := json.Marshal(content)

	_, err = openFile.Write(text)
	goutils.CheckError(err)

	fmt.Println("-->  已写入到sina.json文件中")
}
