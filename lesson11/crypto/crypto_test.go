package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

//单测函数,固定格式
func TestCrypto(t *testing.T) {
	key := "AVSADASFASFASD"
	memfile := new(bytes.Buffer)       //创建一个内存文件
	w := NewCryptoWriter(memfile, key) //写入到内存文件中
	w.Write([]byte("hello"))

	r := NewCryptoReader(memfile, key)
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	if string(buf[:n]) != "hello" {
		t.Errorf("not equal:%s, %s", buf[:n], "hello")
	}

}

//基准测试,固定格式
func BenchmarkCrypto(b *testing.B) {
	buf := []byte(strings.Repeat("a", 1024))

	w := NewCryptoWriter(ioutil.Discard, "123466")
	for i := 0; i < b.N; i++ {
		n, _ := w.Write(buf)
		b.SetBytes(int64(n))
	}
}
