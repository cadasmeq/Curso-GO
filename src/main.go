package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	// Libera la gorutine del wait group
	defer wg.Done()
	println(text)
}

func main() {
	// Declarar wait group
	// El cual va acumulando un grupo de gorutines
	// y las va liberando poco a poco
	var wg sync.WaitGroup

	fmt.Println("hello")

	// Agregar una gorutine al waitgroup
	wg.Add(2)
	go say("world", &wg)
	go say("again", &wg)

	// Indica a la gorutine principal (wait) que espere
	// hasta que el waitgroup finalice con sus gorutines
	wg.Wait()

	// Función anonima que tiene vida solo dentro de si misma
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
