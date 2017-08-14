package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	n, err := strconv.Atoi(s) //字符串转换为int
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(n)
	}

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n)
	}
	aa := "33"
	switch aa {
	case "1":
		fmt.Println("1")
	case "2":
		fmt.Println("2")
	default:
		fmt.Println(aa)

	}
}
