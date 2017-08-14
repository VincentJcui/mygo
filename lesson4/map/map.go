package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int) //创建一个map,key是字符串,value是int
	ages["a"] = 1
	ages["b"] = 2
	fmt.Println(ages)
	fmt.Println(ages["a"])

	// or

	ages1 := map[string]int{ //创建一个map, key是字符串, value是int  并直接赋值
		"a": 1,
		"b": 2,
	}
	fmt.Println(ages1)

	//判断key是否存在
	c, ok := ages["b"]
	if ok {
		fmt.Println("b", c)
	} else {
		fmt.Println("not found")
	}

	//另外一种写法
	if c, ok := ages1["c"]; ok {
		fmt.Println("c", c)
	}
	//添加一个元素
	ages1["c"] = 0
	fmt.Println(ages1)
	//删除一个元素
	delete(ages, "c")
	fmt.Println(ages)
	delete(ages1, "c")
	fmt.Println(ages1)

	//遍历
	for k, v := range ages1 { //遍历key, value, 无序的
		fmt.Println(k, v)
	}
	for k := range ages1 { //遍历key, 无序的
		fmt.Println(k)
	}
}
