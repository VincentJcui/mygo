package main

import (
	"fmt"
	"math"
)

//求一条路径上的所有点的长度

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) GetX() float64 {
	return p.X
}

func Distance(path []Point) float64 {
	var sums float64
	for i := 0; i < len(path)-1; i++ {
		sums += path[i].Distance(path[i+1])
	}
	return sums
}

type Path []Point

func (path Path) Distance() float64 {
	return 0
}

func (path Path) LenPoint() int {
	return len(path)
}

//自定义类型,和方法. type 声明一个类型
type Mystring string

func (s Mystring) substr(i, j int) string {
	return string(s[i:j])
}

func main() {
	path := []Point{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(Distance(path))
	fmt.Println(Path.LenPoint(path))

	//自定义类型和方法
	s1 := Mystring("hello")
	fmt.Println(s1.substr(2, 4))

}
