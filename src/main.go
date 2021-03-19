package main

import "fmt"

func main() {
	helloMessage := "hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de Dato
	fmt.Printf("helloMessage : %T\n", helloMessage)
	fmt.Printf("cursos : %T", cursos)
}
