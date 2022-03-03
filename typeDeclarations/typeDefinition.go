package main

import (
	"fmt"
	"math"
)

// Main types
type Point struct{ i, j float64 }
type Median Point // defines a new distinct type

func main() {
	p1 := Point{2, 3}
	P2 := Point{4, 9}
	fmt.Println(p1.distance(), P2.distance())

	mid := Median{3, 6}
	// fmt.Println(mid.distance()) // new distinct type wont inherit methods
	fmt.Println(mid)

}

// methods
func (p *Point) distance() float64 {
	x := math.Pow(p.i, 2) + math.Pow(p.j, 2)
	return math.Sqrt(x)
}
