package main

import (
	"errors"
	"fmt"
	"strings"
)

func toupper(s string) string {
	return strings.Map(func(i rune) rune {
		return i - ('a' - 'A')
	}, s)
}

func toslow(r rune) rune {
	fmt.Println("%c\n", r)
	return r
}

//闭包
func addn(n int) func(int) int {
	return func(m int) int {
		return m + n
	}

}

func iter(s []int) func() (int, error) {
	var i = 0
	return func() (int, error) {
		if i >= len(s) {
			return 0, errors.New("end")
		}
		n := s[i]
		i += 1
		return n, nil
	}

}

func main() {
	fmt.Println(toupper("hello"))

	s := strings.Map(func(r rune) rune {
		return r - 32
	}, "hello")
	fmt.Println(s)

	//闭包
	f := addn(3)
	fmt.Println(f(2))
	fmt.Println(f(11))

	//iter
	d := iter([]int{1, 2, 3})
	for {
		n, err := d()
		if err != nil {
			break
		}
		fmt.Println(n)
	}
}
