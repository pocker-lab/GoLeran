package GoPing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	// 执行ping baidu的命令, 命令不会结束
	IP := "www.baidu.com"

	_ = Command(IP)
}

// Command 函数接受一个字符串参数 cmd，并执行它作为一个命令
func Command(ip string) error {
	cmd := fmt.Sprintf("ping %v -t -l 100", ip)
	// windows 下使用 cmd 命令行
	c := exec.Command("cmd", "/C", cmd)
	// mac 或 linux 下使用 bash 命令行
	//c := exec.Command("bash", "-c", cmd)
	// --------->
	// 获取命令的标准输出管道
	stdout, err := c.StdoutPipe()
	// 如果出错，返回错误
	if err != nil {
		return err
	}
	// 定义一个同步等待组，用于等待协程结束
	var wg sync.WaitGroup
	// 增加一个等待计数
	wg.Add(1)
	// 启动一个协程，异步读取命令的输出
	go func() {
		// 协程结束时，减少等待计数
		defer wg.Done()
		// 创建一个缓冲读取器，从标准输出管道读取数据
		reader := bufio.NewReader(stdout)
		// 循环读取数据，直到出错或到达文件末尾
		for {
			// 按行读取数据，返回字符串和错误值
			readString, err := reader.ReadString('\n')
			// 如果出错或到达文件末尾，返回并结束协程
			if err != nil || err == io.EOF {
				return
			}
			// 将字节切片转换为字符串（GB18030转换UTF8）
			byte2String := ConvertByte2String([]byte(readString))
			// 打印字符串到标准输出（控制台）
			//fmt.Print(byte2String)
			// 把时间和输出拼接
			now := time.Now()
			//fmt.Printf("%s\n", now.Format("2006-01-02")) // 使用自定义格式打印日期，如：2022-01-14
			//fmt.Printf("%s\n", now.Format("15:04:05"))   // 使用自定义格式打印时分秒，如：16:13:49
			str := fmt.Sprintf("%v %v %v", now.Format("2006-01-02"), now.Format("15:04:05"), byte2String)
			fmt.Print(str)
			WriteFile(ip, str)
			time.Sleep(1 * time.Second)
		}
	}()
	// 启动命令并等待它完成
	err = c.Start()
	// 等待所有协程结束（这里只有一个）
	wg.Wait()
	// 返回错误值（如果有的话）
	return err
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

// ConvertByte2String 将字节切片转换为字符串（GB18030转换UTF8）
func ConvertByte2String(byte []byte) (str string) {
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
	str = string(decodeBytes)
	return str
}
