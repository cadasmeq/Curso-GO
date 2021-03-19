package main

import "fmt"

func main() {
	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println(areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println(result)

	// Resta
	result = y - x
	fmt.Println(result)

	// Multiplicación
	result = x * y
	fmt.Println(result)

	// División
	result = y / x
	fmt.Println(result)

	// Modulo
	result = y % x
	fmt.Println(result)

	// Incremental
	x++
	fmt.Println(x)

	// Decremental
	y--
	fmt.Println(y)

}
