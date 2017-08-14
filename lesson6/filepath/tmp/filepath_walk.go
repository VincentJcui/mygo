package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//遍历文件数

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(path)
		}
		fmt.Println(path)
		return nil
	})
}
