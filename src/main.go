package main

import "fmt"

func inChannel(text string, c chan<- string) {
	c <- text
}

func outChannel(text string, c <-chan string) {
	text = <-c
	fmt.Println(text)
}

func main() {
	c := make(chan string, 1)

	go inChannel("hello", c)
	fmt.Println(<-c)

	fmt.Println("world")

}
