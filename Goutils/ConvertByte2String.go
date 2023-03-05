package Goutils

import "golang.org/x/text/encoding/simplifiedchinese"

// ConvertByte2String 将字节切片转换为字符串（GB18030转换UTF8）
func ConvertByte2String(byte []byte) (str string) {
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
	str = string(decodeBytes)
	return str
}
