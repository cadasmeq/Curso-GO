package main

import "fmt"

func Message(text string, c chan string) {
	c <- text
}

func main() {
	// manejará 2 datos de tipo strings
	c := make(chan string, 2)

	c <- "mensaje 1"
	c <- "mensaje 2"

	// len -> Cuantas corutines hay en el channel
	// cap -> Capacidad máxima que puede almacenar el channel
	fmt.Println(len(c), cap(c))

	// Range y close
	// Cerramos canal
	close(c)

	// c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go Message("este es el mensaje 2", email2)
	go Message("este es el mensaje 1", email1)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2: ", m2)
		}
	}

}
