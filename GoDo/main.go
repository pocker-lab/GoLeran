package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Status int `json:"status"`
	Msg    int `json:"msg"`
}

func main() {
	msg := "{\"status\":200, \"msg\":18}"
	var data Data
	if err := json.Unmarshal([]byte(msg), &data); err == nil {
		fmt.Println(data.Status, data.Msg)
	} else {
		fmt.Println(err)
	}
	var da Data
	da.Status = 188
	fmt.Printf("da.Status: %v\n", da.Status)

}
