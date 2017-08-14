package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x  byte
		x1 int
		x2 int32
		x3 int64
		x4 uint
		x5 uint32
		x6 uint64
		x7 int8  = 127
		x8 uint8 = 255
	)
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
	for i := 0; i < 7; i++ {
		if i == 0 {
			fmt.Println(unsafe.Sizeof(x))
			continue
		}
		fmt.Println(unsafe.Sizeof(string(x) + string(i)))
	}
	fmt.Println(unsafe.Sizeof(x1))
	fmt.Println(unsafe.Sizeof(x7), x7, x7+1)
	fmt.Println(unsafe.Sizeof(x8), x8, x8+1)
}
