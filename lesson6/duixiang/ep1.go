package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

//下面这里传入的值或者指针都可以跑
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))
}
