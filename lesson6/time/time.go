package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.After(3))
	var m, n time.Duration //定义一个时间变量
	n = 2 * time.Second
	m = 3*time.Hour + 10*time.Minute + 45*time.Second
	/*
		time.Hour  小时
		time.Second  秒
		time.Minute  分钟
	*/
	//time.Sleep(n)
	fmt.Println(int64(n))
	fmt.Println(n.String())
	fmt.Println(n.Seconds())
	fmt.Println(n.Minutes())
	//
	fmt.Println(m.String())
	fmt.Println(m.Seconds())
	fmt.Println(m.Minutes())

	//定义t  = 当前时间
	t := time.Now()
	//获取一小时之后的时间
	after := t.Add(time.Hour)
	fmt.Println(after)
	//
	before := t.Add(-1 * time.Hour)
	fmt.Println(before)
	//计算时间差,sub代表t1 - t 的时间
	t1 := t.Add(time.Hour)
	fmt.Println(t1.Sub(t))

}
