package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int //统计字节数

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type LineCounter int //统计行数

func (d *LineCounter) Write(p []byte) (int, error) {
	*d = LineCounter(bytes.Count(p, []byte("\n")))
	return len(p), nil
}

/*
另外的写法

type LineCounter struct {
	Sum int
}

func (d *LineCounter) Write(p []byte) (int , error)  {
	for _, b := range p{
		if b == '\n'{
			d.Sum++
		}
	}
}


*/

func main() {
	b := new(ByteCounter)
	d := new(LineCounter)
	//io聚合
	w := io.MultiWriter(b, d)
	io.Copy(w, os.Stdin)
	fmt.Println(*b)
	fmt.Println(*d)

	//内存buf
	buf := new(bytes.Buffer)
	buf.WriteString(`
	hello gopher
	123345
	main new
	`)
	ww := io.MultiWriter(b, d)
	io.Copy(ww, buf)
	fmt.Println(*b)
	fmt.Println(*d)
}
