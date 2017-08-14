package main

import (
	"fmt"
)

//100以内的斐波那契数列之和
//1,1,2,3,5,8,13,21...

func main() {
	var x = 1
	var y = 1
	var sum = x + y
	fmt.Printf("%d \n%d \n", x, y)
	for y < 100 {
		x, y = y, x+y
		if y >= 100 {
			break
		}
		sum += y
		fmt.Println(y)
	}
	fmt.Println("sum =", sum)
}
