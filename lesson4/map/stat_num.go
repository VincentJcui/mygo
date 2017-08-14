package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func stat_num(s []string) {
	nums := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		n, ok := nums[s[i]]
		if ok {
			nums[s[i]] = n + 1
		} else {
			nums[s[i]] = 1
		}
	}
	for k, v := range nums {
		fmt.Println(k, ":", v)
	}
}

func main() {
	var s = [...]string{
		"abc",
		"dd",
		"ee",
		"abc",
		"ee",
	}
	//fmt.Println(s)
	stat_num(s[0:])
	/*  原始方法一
	if len(os.Args) < 1{
		fmt.Println("not found files")
		return
	}
	w, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(w)
	re := ""
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		re += line
	}
	//fmt.Println(re)
	stat_num(strings.Fields(re))   //strings.Fields 是将字符按单词分隔  具体参考:  http://godoc.org/strings#example-Fields
	*/

	//方法二  ioutil 只针对小文件
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	stat_num(strings.Fields(string(content)))
}
