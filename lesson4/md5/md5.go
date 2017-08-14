package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func md5s(s string) {
	data := []byte(s)
	md5sum := md5.Sum(data)
	fmt.Printf("%v\n%x\n", md5sum, md5sum)
}

func main() {
	//fmt.Printf("%x", 255) //打印255的16进制

	//data := []byte("hello")
	//md5sum := md5.Sum(data)
	//fmt.Printf("%x\n", md5sum)
	if len(os.Args) < 2 {
		fmt.Println("args input error")
		return
	}
	files := os.Args[1]
	w, err := os.Open(files)
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
	md5s(re)

}
