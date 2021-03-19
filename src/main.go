package main

import "fmt"

func Printer(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Funciones
	Printer("Hola Mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(5)
	fmt.Println(value)

	value1, value2 := doubleReturn(5)
	fmt.Println(value1, value2)

	value1, _ = doubleReturn(5)
	fmt.Println(value1, value2)
}
