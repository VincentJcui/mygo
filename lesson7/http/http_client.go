package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func cleanUrls(url string, urls []string) []string {
	/*
		清洗url的形式

		http://xxx.com/a.jpg
		//xx.com/a.jpg
		/ststic/a.jpg
		a.jpg
	*/
	return []string{}
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("http error code:%s", resp.Status)
		return nil, err
	}
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := (s.Attr("src"))
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}
	})
	return urls, err
}

func main() {
	url := "http://daily.zhihu.com/"
	/*
		http://59.110.12.72:7070/golang-spider/img.html
	*/
	//url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}
