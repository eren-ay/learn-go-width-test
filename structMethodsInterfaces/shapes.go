package shape

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// simple func
func rectPerimeter(rectangle Rectangle) float64 {
	perimeter := 2 * (rectangle.Width + rectangle.Height)
	return perimeter
}

func rectArea(rectangle Rectangle) float64 {
	area := rectangle.Width * rectangle.Height
	return area
}

// this func calling like Circle.Area()= this func belongs to circle struct
func (c Circle) Area() float64 {
	area := math.Pi * c.Radius * c.Radius
	return area
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
