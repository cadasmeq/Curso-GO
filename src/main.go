package main

import "fmt"

func calculateArea(s shape2D) {
	fmt.Println("Area: ", s.computeArea())
}

type shape2D interface {
	computeArea() float64
}

type Square struct{ base float64 }
type Rectangle struct {
	height float64
	width  float64
}

func (s Square) computeArea() float64 {
	return s.base * s.base
}

func (r Rectangle) computeArea() float64 {
	return r.height * r.width
}

func main() {
	square := Square{base: 5}
	rectangle := Rectangle{height: 5, width: 3}

	calculateArea(square)
	calculateArea(rectangle)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, true}
	fmt.Println(myInterface...)

}
