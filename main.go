package main

import (
	"fmt"
	"go_do/data_center"
)

type Data struct {
	TotalNum string `json:"total_num"`
	Symbol   string `json:"symbol"`
}

func main() {

	sinaBody := data_center.SinaGetRequest(1)
	sinaBody2 := data_center.SinaGetRequest(5)
	//
	fmt.Println(sinaBody)
	fmt.Println("-----------------")
	fmt.Println(sinaBody2)
	////fmt.Printf("%T", sinaBody)
	//
	//fmt.Print("\n--------------------------------")
	//re := regexp.MustCompile(`(?m)/\*.*\*/`)
	//
	//sinaBody1 := re.ReplaceAllString(sinaBody, "")
	//fmt.Print(sinaBody1)

	test()

}
