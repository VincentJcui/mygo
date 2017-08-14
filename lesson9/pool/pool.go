package main

import (
	"fmt"
	"log"
	"net/http"

	"sync"
)

//给定一个url返回url,打印url和url的status
//www.baidu.com 200 ok
func printurl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}
func work(ch chan string, wg *sync.WaitGroup) {
	//url := <-ch
	//printurl(url)
	//方法一
	//for{
	//	if url,ok := <-ch;ok{
	//		printurl(url)
	//	}else{
	//		break
	//	}
	//}
	//方法二
	for url := range ch {
		printurl(url)
	}
	wg.Done()
}

//channel特色
//1.只要不close可以永远发送数据和接收数据
//2.如果channel里面没有数据,接收方会阻塞
//3.如果没有人正在等待channel数据,发送方会阻塞
//4.从一个close的channel取数据永远不会阻塞,同时获取的是默认值

//主程成启动work协程,同时传递一个channel
//主协程channek里面发送url
//work协程从channel里面获取url,之后调用printurl打印url

//启动3个协程
//主协程同时向channel发送多个url,发送完毕之后关闭chann
//work协程从chann里面获取url,之后调用printurl打印url
//work协程不停重复第三条,直到channel关闭

//创建一个WaitGroup
//调用Add
//调用Wait等待work协程结束
func main() {
	//单url
	//ch := make(chan string)
	//go work(ch)
	//url := "http://www.baidu.com"
	//ch <- url
	//time.Sleep(3*time.Second)

	//
	wg := new(sync.WaitGroup)
	//wg.Add(3)   //一般起几个协程,Add几个协程,可以改成如下几个方式
	urls := []string{"http://www.baidu.com", "http://www.qq.com", "http://www.163.com"}
	ch := make(chan string)
	for i := 0; i < 3; i++ {
		//或者这样,每起一个协程,Add一个
		wg.Add(1)
		go work(ch, wg)
	}

	for _, url := range urls {
		ch <- url
	}
	//time.Sleep(3*time.Second)
	close(ch)
	wg.Wait()
}
