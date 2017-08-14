package main

import "fmt"

func test() {
	var a, c int
	a = 10
	c = 3
	fmt.Println(a + c)
	fmt.Println(a - c)
	fmt.Println(a * c)
	fmt.Println(a / c)
	fmt.Println(a % c)

	if (a > c && c > 1) || c > 20 {
		fmt.Println(c + c + c)
	}

}

func main() {
	var b bool
	b = true
	b = false
	b = ("hellp" == "world")
	if b {
		fmt.Println("相等")
	} else {
		fmt.Println("不等")
	}

	fmt.Println(3 / 2.0)
	fmt.Printf("%.2f \n", 3.0/2)
	test()
	var d bool
	if b == d || b != d {
		fmt.Println("b=d")
	}

}
