package main

type triangle struct {
	base, height float64
}

func (t *triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func main() {
	// Define a method on the triangle called Area
	// which returns the area covered by the triangle.
}
