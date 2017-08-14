package main

// 解压缩tar的文件并打印文件名字
import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	tr := tar.NewReader(os.Stdin)
	for {
		hdr, err := tr.Next() //获取文件头信息
		if err != nil {
			return
		}
		fmt.Println(hdr.Name)
		info := hdr.FileInfo()
		//io.Copy(ioutil.Discard, tr) //将内容丢弃,等同于>/dev/null
		if info.IsDir() {
			os.Mkdir(hdr.Name, 0755)
			continue
		}
		f, err := os.Create(hdr.Name)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(f, tr)
		os.Chmod(hdr.Name, info.Mode())
		f.Close()
	}
}
