package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a { //遍历下标和值
		fmt.Printf("%d %d \n", i, v)
	}
	for _, v := range a { //遍历值,忽略下标
		fmt.Printf("%d\n", v)
	}
	for i := range a { //遍历下标,忽略值
		fmt.Printf("%d\n", i)
	}

	var q [3]int = [3]int{1, 2, 3} //定义数组规定包含3个元素,并初始化值为1,2,3
	var r [3]int = [3]int{1, 2}    //定义数组规定包含3个元素,并初始化前两个值1,2
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(r[2])

	q1 := [...]int{1, 2, 3} //定义一个无限数组,并初始化三个值为1,2,3
	fmt.Println(q1)

	q2 := [...]int{4: 2, 10: -1} //定义一个无限数组,并初始化下标为4的值为2, 下标为10的值为-1
	fmt.Println(q2)
	fmt.Println(len(q2))

	var d [3]int
	d = q
	fmt.Println(d)
	fmt.Println(&d[0], &q[0])     //& 打印地址
	fmt.Println(unsafe.Sizeof(q)) //打印占用的字节,内存中的字节
	fmt.Println(unsafe.Sizeof(d))

}
