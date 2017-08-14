package read

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Read ,按块读取,效率较低  裸读取,很少使用
	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	fmt.Println(buf[:n])

	//bufio .加上buffer的读取, 高效的读取方式
	r := bufio.NewReader(f)
	r.ReadByte() //按字节读取
	r.Read(buf)

	//按行读取,按分隔符读取
	r1 := bufio.NewScanner(f)

	//小文件一次性读取
	ioutil.ReadFile("a.txt")
	ioutil.ReadAll(f)

	//神器, io.copy 操作类文件的高级方法
	io.Copy()

}
