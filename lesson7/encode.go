package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	//首字母大写,表示方法是共用的
	Name string
	id   int
	//首字母小写,表示不希望别人使用的方法
	//name string
	//id int
}

func (s *Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func main() {
	s := Student{
		Name: "binggan",
		id:   1,
	}
	buf, err := json.Marshal(s)
	//buf1, err1 := json.Marshal(s1)
	if err != nil {
		log.Fatal(err)
	}
	//if err1 != nil{
	//	log.Fatal(err1)
	//}
	fmt.Println(string(buf))
	//fmt.Println(string(buf1))
}
