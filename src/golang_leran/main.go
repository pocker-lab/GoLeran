package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[:5]  //1-5
	s2 := arr[3:8] //4-8
	s3 := arr[5:]  //6-10
	s4 := arr[:]   //1-10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
