package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second) //每秒触发
	cnt := 0
	for _ = range timer.C {
		cnt++

		if cnt > 10 {
			timer.Stop()
			return
		}
		fmt.Println("hello", cnt)

	}
	c := time.After(time.Second * 3)
	<-c
	fmt.Println("done")
}
