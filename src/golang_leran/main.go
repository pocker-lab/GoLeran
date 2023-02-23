package main

import "fmt"

func main() {
	letter := "A"
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音")
	case "N", "M":
		fmt.Println("M或N")
	default:
		fmt.Println("其他")
	}
}
