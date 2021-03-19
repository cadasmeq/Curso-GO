package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("hola")
	fmt.Println("mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 6 {
			fmt.Println("break")
		}
	}

}
