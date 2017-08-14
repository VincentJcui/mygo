package main

import (
	"fmt"
)

func main() {
	fmt.Println()          //打印并换行
	fmt.Printf("%v \n", 1) //格式化
	//fmt.Fprintf(f, "%v", 1)       //f代表写入到文件

	s := fmt.Sprintf("http://%s/%s", "www.baidu.com", "/pic")
	fmt.Println(s)

	var x string
	var y int
	//var line string
	fmt.Scanf("%s %d", &x, &y)
	fmt.Println(x, y)
	for {
		fmt.Print(">")

		//fmt.Printf("line:#%s#", line)
		fmt.Scan(&x, &y)
		if x == "stop" {
			break
		}
		fmt.Println(x, y)
	}
}
