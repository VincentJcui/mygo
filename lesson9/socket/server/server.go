package main

import (
	"log"
	"net"
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
	//conn.Write([]byte("hello golang\n"))
	conn.Write([]byte(message))
	//time.Sleep(time.Second * 10)
	conn.Close()
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
