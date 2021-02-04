package main

import "math"

type shape interface {
	getArea() float64
}

type Triangle struct {
	width  int
	height int
}

type Square struct {
	sideLength int
}

func (t Triangle) getArea() float64 {
	return 0.5 * float64(t.width) * float64(t.height)
}

func (s Square) getArea() float64 {
	return math.Pow(float64(s.sideLength), 2)
}
