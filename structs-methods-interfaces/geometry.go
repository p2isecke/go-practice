package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
}

// Rectangle has a method called Area that returns float64
type Rectangle struct {
	length float64
	width float64
}

// area of a rectangle: l * w
func (r Rectangle) Area() float64 {
	return r.length * r.width
}


type Circle struct {
	radius float64
}

// Circle has a method called Area that returns float64
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base float64
	height float64
}

// area of triangle: 1/2 * base * perpendicular height
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// perimeter of a rectangle: 2*(l+w)
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.width )
}
