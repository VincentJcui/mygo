package main

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

//把b里面的数据进行加密,之后写入到w.w里面
//调用w.w.Write方进行写入
func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b)
	return w.w.Write(buf)
}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}
}

func (r *CryptoReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}

func Crypto(conn net.Conn, r io.Reader, key string) {
	//创建cipher
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Fatal(err)
	}
	//创建buf
	buf := make([]byte, 4096)
	for {
		// 从r里面读取数据到buf
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		// 加密buf
		cipher.XORKeyStream(buf[:n], buf[:n])
		// 把buf写入到w里面
		conn.Write(buf[:n])
	}

}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Print("Error:", err)
		return
	}
	//key := GetRandomString(10)
	key := "ABCDefgQWERasdf"
	//fmt.Println(key)
	//remote.Write([]byte(key))
	//加密的过程

	//r_buf := bufio.NewReader(conn)
	//Crypto(remote, r_buf, key)
	//fmt.Println(server)

	//向服务器端发送数据

	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 接收客户端的数据,发送到remote,直到conn的EOF
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	//解密的过程

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
		log.Print(conn.RemoteAddr())
		if err != nil {
			log.Print(err)
		}
		// 参考页面 http://www.jianshu.com/p/172810a70fad
		go handleConn(conn)
	}
}
