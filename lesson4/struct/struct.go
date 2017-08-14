package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	var s Student //声明一个结构体
	s.Id = 1
	s.Name = "jcui"
	fmt.Println(s)

	s.Id = 2
	s.Name = "wuke"
	fmt.Println(s)

	//另外一种声明方式
	s1 := Student{

		Id:   1,
		Name: "haha",
	}
	fmt.Println(s1)

	s1 = s
	fmt.Println(s1)

	//结构体指针

	var p *Student
	p = &s1
	p.Id = 3
	fmt.Println(s1)

	var p1 *int
	p1 = &s1.Id
	*p1 = 4
	fmt.Println(s1)

	var arr [3]Student //声明一个数组结构体
	var ss []Student
	var m map[string]Student //声明一个map结构体
	fmt.Println(arr, ss, m)
}
