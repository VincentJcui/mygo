package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func subtraction(x, y int) int {
	return x - y
}

func multiplication(x, y int) int {
	return x * y
}

func division(x, y int) int {
	return x / y
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("args error")
		//os.Exit(0)
	}
	z := os.Args[2]
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	//fmt.Println(x, y, z)
	if string(z) == "+" {
		fmt.Println(add(x, y))
	} else if string(z) == "-" {
		fmt.Println(subtraction(x, y))
	} else if string(z) == "*" {
		fmt.Println(multiplication(x, y))
	} else if string(z) == "\\" {
		fmt.Println(division(x, y))
	} else {
		fmt.Println("args error error")
	}

}
