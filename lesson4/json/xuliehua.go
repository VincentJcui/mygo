package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := Student{
		Id:   2,
		Name: "jcui",
	}

	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal("marshal error: %s", err)
	}
	fmt.Println(string(buf))
}
