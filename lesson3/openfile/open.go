package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//f, err := os.Create("a.txt")  //如果文件不存在则创建,如果文件存在则清空
	//if err != nil {
	//	log.Fatal(err)
	//}
	//f.WriteString("hello\n")
	//f.Close()

	w, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//w.WriteString("hello\n")
	//w.Seek(1, os.SEEK_SET)
	//w.WriteString("$$")
	//w.Seek(0, os.SEEK_SET)
	//w.WriteString("!@#$")

	//buf := make([]byte, 1024)
	//w.Read(buf)
	//fmt.Println(string(buf))
	r := bufio.NewReader(w)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	w.Close()

}
