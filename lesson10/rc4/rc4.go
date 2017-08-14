package main

import (
	"crypto/md5"
	"crypto/rc4"
	"flag"
	"io"
	"log"
	"os"
)

var (
	key = flag.String("-k", "123456", "secort string")
)

func crypto(w io.Writer, r io.Reader, key string) {
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
		w.Write(buf[:n])
	}
}

func main() {
	/*
		key := "123456"
		md5sum := md5.Sum([]byte(key))
		//加密过程
		cipher, err := rc4.NewCipher([]byte(md5sum[:]))
		if err != nil {
			log.Fatal(err)
		}

		buf := []byte("what a fuck you")
		cipher.XORKeyStream(buf, buf)
		log.Print(string(buf))

		//解密过程
		{
			cipher, err := rc4.NewCipher([]byte(md5sum[:]))
			if err != nil {
				log.Fatal(err)
			}
			cipher.XORKeyStream(buf, buf)
			log.Print(string(buf))
		}
	*/
	flag.Parse()
	crypto(os.Stdout, os.Stdin, *key)

}
