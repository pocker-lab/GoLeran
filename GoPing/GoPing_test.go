package GoPing

import "testing"

func TestCommand(t *testing.T) {
	IP := "www.baidu.com"
	_ = Command(IP)
}
