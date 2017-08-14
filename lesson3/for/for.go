package main

import (
	"fmt"
	"os"
)

func main() {
	//方式一
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	//方式二
	i := 5
	for i < 7 {
		fmt.Println(i)
		i += 1
	}
	//方式三
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	//方式四
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	s := "hello"
	for i, arg := range s {
		fmt.Println(i, arg)
	}
}
