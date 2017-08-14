package main

//强类型
//类型转换需要显示进行
//字符串跟数字需要借助函数转换
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var n int
	var f float32
	var z float32
	n = 10
	// f = n / 3     注意左右两侧的数据类型是不一样的,所以程序会报错
	f = float32(n) / 3
	z = float32(n / 3)
	fmt.Println(f, z, n)
	n = int(f) * 10
	fmt.Println(f, n)
	n = int(f * 10)
	fmt.Println(f, n)

	var n1 int64
	var n2 int8
	n1 = 1024004
	n2 = int8(n1)
	fmt.Println(n2)

	n1 = 1024129
	n2 = int8(n1)
	fmt.Println(n2)

	var s string
	s = strconv.Itoa(n)
	fmt.Println(s)
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err) //如果转换的字符不是数字,则会抛出错误
	}
	fmt.Println(n)

	//模拟随机字符串
	var x int64
	var s1 string
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 36)
	s1 = strconv.FormatInt(x, 10)
	fmt.Println(s)
	fmt.Println(s1)
	//世界上只有1和0两种人,一种懂二进制,一种不懂二进制
}
