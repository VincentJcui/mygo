package main

import (
	"fmt"
)

//多返回值
func swap(x, y string) (string, string) {
	return y, x
}

//可变参数,注意这里的点点点
func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

//命名返回值
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}

//递归,实现斐波纳挈数列
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

/*
通项公式 : a(n) = 2 * a(n-1) + n -1
第一项: a(1) = 2
求第10项
*/
func fib1(n int) int {
	if n <= 1 {
		return 2
	}
	return (2*fib1(n-1) + (n - 1))
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(171))
	fmt.Println(sum(1, 2, 3))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(s...)) //注意这里的点点点
	fmt.Println(fib(2))
	fmt.Println(fib1(10))
}
