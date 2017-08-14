package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {

	//s := os.Args[1]
	s := "http://59.110.12.72:7070/golang-spider/img.html"
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("scheme", u.Scheme)
	fmt.Println("host", u.Host)
	fmt.Println("path", u.Path)
	fmt.Println("queryString", u.RawQuery)
	fmt.Println("user", u.User)
	fmt.Println("xxx", u.Fragment)
}
