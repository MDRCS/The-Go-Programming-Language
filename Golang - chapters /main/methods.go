package main

import (
	"fmt"
	"math"
)

type point struct {
	X,Y float64
}

func (p point) distance(q point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func main() {
	var p = point{1.1,2.1}
	var q = point{3.1,6.1}

	var dist = p.distance(q)

	fmt.Printf("%.2f",dist)

}