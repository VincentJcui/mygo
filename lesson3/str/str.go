package main

import (
	"fmt"
)

func toupper(s string) string {
	array := []byte(s)
	for i := 0; i < len(array); i++ {
		if array[i] >= 'a' && array[i] <= 'z' {
			array[i] = array[i] - 32
		} else if array[i] <= 'Z' && array[i] >= 'A' {
			array[i] = array[i] + ('a' - 'A')
		}
	}
	return string(array)
}

func main() {
	str1 := "hello"
	doc := `
	你好
	我的世界
	` //注意多行打印使用反引号,等同于python的 '''字符串'''
	ff := "\"\"" //注意双引号中要打印双引号需要在被打印前加斜杠\
	test3 := "\"你好我的世界?\",he said"
	fmt.Println(str1)
	fmt.Println(doc)
	fmt.Println(ff)
	fmt.Println(test3)
	//相加
	s1 := "hello" + "world"
	// 取字符
	var c1 byte
	fmt.Println(0, len(s1)-1)
	c1 = s1[0]
	c2 := s1[0]
	c3 := s1[0:1]
	//切片s
	s2 := s1[0:3]
	s3 := s1[:]

	fmt.Println(s1)
	fmt.Println(c1)
	fmt.Printf("%d, %c ,%b \n", c1, c1, c1) //%d  代表十进制数值   %c 代表unicode字符  %b 代表二进制
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(s2)
	fmt.Println(s3)

	var b byte
	for b = 0; b < 177; b++ {
		fmt.Printf("%d,%c \n", b, b)
	}

	fmt.Println(0xA) //十六进制是以0x开头

	//修改字符串
	//将s1强制转换
	array := []byte(s1)
	fmt.Println(array) //输出一个数组,asic码 [104 101 108 108 111 119 111 114 108 100]

	array[0] = 'H'
	for i := 0; i < len(array)-1; i++ {
		if array[i] == 'w' {
			array[i] = 'W'
		}
	}
	s1 = string(array)
	fmt.Println(s1)
	fmt.Println('a' + ('H' - 'h'))
	fmt.Printf("%c \n", 'a'+('H'-'h'))
	fmt.Println('c' + ('H' - 'h'))
	fmt.Printf("%c \n", 'c'+('H'-'h'))
	fmt.Println('g' + ('T' - 't'))
	fmt.Printf("%c \n", 'g'+('T'-'t'))

	fmt.Println(toupper("asdasdzZ"))
	fmt.Println(toupper("ZASDFGH"))

}

//  linux xxd str.go  可以查看str.go的十六进制
