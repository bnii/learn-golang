package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimiter(rec Rectangle) (perimiter float64) {
	perimiter = 2*rec.Height + 2*rec.Width
	return
}

func Area(rec Rectangle) (area float64) {
	area = rec.Width * rec.Height
	return
}
