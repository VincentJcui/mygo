package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	go say("world")
	say("hello")

	//sleep排序
	s := []int{2, 7, 1, 6, 4}
	for _, n := range s {
		go func(n int) { //go起尼玛函数
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(n)
	}
	time.Sleep(10 * time.Second)

	//channel
	p := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) //声明一个channel
	go sum(p[:len(p)/2], c)
	go sum(p[len(p)/2:], c)
	x, y := <-c, <-c //receive form c (channel)
	fmt.Println(x, y, x+y)
}
