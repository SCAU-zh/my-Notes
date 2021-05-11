package learnGoWithTests

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base float64
	height float64
}

type Geometry interface {
	Perimeter() float64
	Area() float64
}

func (r *Rectangle) Perimeter() float64{
	return (r.Width + r.Height) * 2
}

func (r *Rectangle) Area() float64{
	return r.Width * r.Height
}

func (r *Circle) Perimeter() float64{
	return 3.0/4.0 * math.Pi * r.Radius
}

func (r *Circle) Area() float64{
	return math.Pi * r.Radius *r.Radius
}

func (r *Triangle) Perimeter() float64{
	return math.Sqrt(math.Pow(r.Base/2, 2) + math.Pow(r.height, 2))*2 + r.Base
}

func (r *Triangle) Area() float64{
	return r.Base * r.height / 2
}

