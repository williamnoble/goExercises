package shapes

import "math"

type Rectangle struct {
	width  float64
	length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

// Perimeter adds the length of all four sides e.g. w + w + h + h
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
