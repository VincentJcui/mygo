package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Students struct {
	Id   int
	Name string
}

func main() {
	str := `{"Id":2 ,"Name":"jcui"}`
	var s Students
	err := json.Unmarshal([]byte(str), &s) //注意这里load到内存中传的是指针
	if err != nil {
		log.Fatal("unmarshal error:%s", err)
	}
	fmt.Println(s)
}
