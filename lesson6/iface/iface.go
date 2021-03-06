package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance2Point(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Distance() float64 {
	return math.Hypot(p.X, p.Y)
}

type Path []*Point

func (p Path) Distance() float64 {
	var sum float64
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].Distance2Point(p[i+1])
	}
	return sum
}

func print(p IDistance) {
	fmt.Println(p.Distance())
}

//func handle(w http.ResponseWriter , r *http.Request)  {
//	fmt.Fprintf(w, "hello \n")
//}
//声明接口
type IInstance interface {
	Instance() float64
}

type IDistance interface {
	Distance() float64
}

func main() {
	//var i IInstance
	//p := Path{{1,2},{3,4}}
	//i = p
	//fmt.Println(i)
	//
	var path Path
	path = make([]*Point, 3)
	p1 := &Point{1, 2}
	p2 := &Point{3, 4}
	p3 := &Point{5, 6}
	path[0] = p1
	path[1] = p2
	path[2] = p3
	var i IDistance
	i = p1
	fmt.Println(i.Distance())
	i = p2
	fmt.Println(i.Distance())
	i = path
	fmt.Println(i.Distance())
	print(path)
	print(p1)

}
