package main

import (
	"bufio"
	"flag"
	"fmt"
	goutils "github.com/pocker-lab/goutils"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

var (
	ip string
)

func main() {
	flag.StringVar(&ip, "ip", "localhost", "ip地址")
	flag.Parse()
	// 拼接命令
	order := fmt.Sprintf("ping %v -t -l 100", ip)

	wdstr, _ := os.Getwd()
	fmt.Printf("ip地址：\t%v\n", ip)
	fmt.Printf("文件地址：\t%v\\%v.txt\n", wdstr, ip)

	c := exec.Command("cmd", "/C", order)
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Printf("报错：%v", err)
		return
	}
	start := time.Now()
	var wg sync.WaitGroup // 定义一个同步等待组，用于等待协程结束
	wg.Add(1)             // 增加一个等待计数
	go func() {           // 启动一个协程，异步读取命令的输出
		defer wg.Done()                   // 协程结束时，减少等待计数
		reader := bufio.NewReader(stdout) // 创建一个缓冲读取器，从标准输出管道读取数据
		for {                             // 循环读取数据，直到出错或到达文件末尾
			readString, err := reader.ReadString('\n') // 按行读取数据，返回字符串和错误值
			if err != nil || err == io.EOF {
				return
			} // 如果出错或到达文件末尾，返回并结束协程
			byte2string := goutils.ConvertByte2String([]byte(readString)) // 转换编码为UTF-8
			// 拼接每行输出的string
			now := time.Now()
			str := fmt.Sprintf("%v %v %v", now.Format("2006-01-02"), now.Format("15:04:05"), byte2string)
			// 将输出写入到文件中
			WriteFile(ip, str)
			// 实时显示文件运行时间
			time.Sleep(time.Second)
			elapsed := time.Since(start)
			hours := int(elapsed.Hours())
			minutes := int(elapsed.Minutes()) % 60
			seconds := int(elapsed.Seconds()) % 60
			fmt.Printf("\r运行时长: %02d:%02d:%02d", hours, minutes, seconds)
		}
	}()

	err = c.Start() // 启动命令并等待它完成
	wg.Wait()       // 等待所有协程结束（这里只有一个）
}

func WriteFile(ip, str string) {
	// 打开或创建一个名为 output.txt 的文件，指定写入模式为追加（os.O_APPEND）和创建（os.O_CREATE），权限为 0666
	file, err := os.OpenFile(fmt.Sprintf("%v.txt", ip), os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err) // 如果出错，打印错误并退出程序
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file) // 延迟关闭文件

	_, err = file.WriteString(str) // 将字符串写入到文件中，并返回写入的字节数和错误值
	if err != nil {
		fmt.Println(err) // 如果出错，打印错误并退出程序
		return
	}
}
