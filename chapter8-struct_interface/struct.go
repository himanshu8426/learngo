package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
func main() {

	//c := Circle{x: 1.0, y: 2.0, r: 5.0}
	//fmt.Println(math.Pi*c.r*c.r, c.x, c.y)

	c := Circle{0, 0, 5}

	fmt.Println(circleArea(c))
}
