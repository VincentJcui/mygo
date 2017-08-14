package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

/*
	监听地址
	接受连接

*/
var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {

	//建立到目标服务器的连接
	var remote net.Conn
	remote, err := net.Dial("tcp", *target)
	defer remote.Close()
	if err != nil {
		log.Print(err)
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 接收客户端的数据,发送到remote,直到conn的EOF
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	//go 接收remote的数据,发送到客户端,直到remote的EOF
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
}

func main() {
	//建立监听
	flag.Parse()
	addr := ":7777"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//接收新的连接
	for {
		//accept new conection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
