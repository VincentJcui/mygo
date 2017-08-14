package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

/*
监听端口
接受新的连接
启动协程
发送接收数据
断开连接
*/

var message = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:03:49 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/
<html>
<body>
<h1 style="color:red">hello golang</h1>
</body>
</html>

`

func handleConn(conn net.Conn) { //主机conn 这里的类型net.Conn
	defer conn.Close()
	//读取客户端需求,获得客户端需要得到的文件
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input \n"))
		return
	}
	cmd := fields[0]
	name := fields[1]
	switch cmd {
	case "GET":
		//打开文件
		r, err := os.Open(name)
		if err != nil {
			log.Print(err)
			return
		}
		defer r.Close()

		//读取内容
		buf, err := ioutil.ReadAll(r)
		////发送内容
		conn.Write(buf)
	case "LS":
		fmt.Println(cmd)
	case "STORE":
		// 从r读取文件直到err为io.EOF
		//创建name文件
		//向文件写入数据
		//往conn写入OK
		//关闭连接和文件
	}

}

func main() {
	addr := ":7777" //监听任意IP的7777端口
	//创建监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

	}

}
