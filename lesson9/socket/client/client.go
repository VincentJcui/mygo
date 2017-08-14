package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := "www.baidu.com:80"
	//拨号,请求连接
	conn, err := net.Dial("tcp", addr)
	//记得关闭
	defer conn.Close()
	//如果没有错误,则表示连接成功
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String()) //打印远端地址及端口
	fmt.Println(conn.LocalAddr().String())  //打印本地地址及建立连接所随机的端口

	//发送数据
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n) //代表发送了18个字节

	//读取返回结果
	//buf := make([]byte, 4096) //创建一个缓冲区
	//n, err = conn.Read(buf)  //n代表读取的行数
	//if err != nil && err != io.EOF{   //读取到结尾,EOF代表对方把连接关闭
	//	log.Fatal(err)
	//}
	//fmt.Println(n, string(buf[:n]))

	//另外一种简介的输出结果
	io.Copy(os.Stdout, conn)

}
