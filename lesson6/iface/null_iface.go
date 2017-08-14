package main

import "fmt"

type Writer interface {
	Writer(b []byte) (int, error)
}

//type I interface {
//
//}

func main() {
	var i interface{}
	var n int
	i = n
	var s string
	i = s
	//var p Point
	//i = p
	fmt.Println(i)

}
