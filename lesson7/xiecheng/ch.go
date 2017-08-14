package main

import (
	"fmt"
)

func pp(s []string, c chan string) {
	res := s[0]
	for i := 1; i < len(s); i++ {
		//time.Sleep(time.Second * time.Duration(i))
		res += s[i]
	}
	c <- res
}

func main() {
	s := []string{"hello", "golang", "c++", "world"}
	c1 := make(chan string)
	c2 := make(chan string)
	go pp(s[:len(s)/2], c1)
	go pp(s[len(s)/2:], c2)
	x, y := <-c1, <-c2
	fmt.Println(x + y)

	//channel的缓冲区大小

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
