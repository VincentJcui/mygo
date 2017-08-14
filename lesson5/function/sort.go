package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := []int{2, 3, 8, 9, 4}
	s1 := []rune("helloworld")
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	fmt.Println(string(s1))

	ss := []Student{}
	ss = append(ss, Student{
		Name: "aa",
		Id:   3,
	})
	ss = append(ss, Student{
		Name: "bb",
		Id:   1,
	})
	ss = append(ss, Student{
		Name: "cc",
		Id:   2,
	})

	//针对结构体排序
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id < ss[j].Id
	})
	fmt.Println(ss)
}
