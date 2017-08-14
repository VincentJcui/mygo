package main

import "fmt"

const (
	PI = 3.1415926
	E  = 2.0
	G  = 9.8
)

const (
	A = iota * 10
	B
	C
)

func main() {
	fmt.Println(A, B, C)
}
