package main

import (
	"log"
	"net/http"
	"sync"
)

func work(ch chan string, wg *sync.WaitGroup) {
	for u := range ch {
		resp, err := http.Get(u)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("%s,%d", u, resp.ContentLength)
	}
	wg.Done()

}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	taskch := make(chan string)
	for i := 0; i < 5; i++ {
		go work(taskch, wg)
	}

	urls := []string{"http://www.baidu.com", "http://www.qq.com"}
	for _, url := range urls {
		taskch <- url
	}
	close(taskch)
	wg.Wait()
}
