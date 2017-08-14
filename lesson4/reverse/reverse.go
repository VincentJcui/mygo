package main

import (
	"fmt"
)

func reverse(s []int) {

}

func main() {
	s := []int{2, 3, 5, 7, 11}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
	reverse(s[1:4])
	fmt.Println(s)

	str := "hello world"
	fmt.Println(str) //要求输出结果"world hello"
}
