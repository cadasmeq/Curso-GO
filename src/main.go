package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	mypc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(mypc)

	mypc.ping()

	fmt.Println(mypc)

	mypc.duplicateRAM()
	fmt.Println(mypc)

	mypc.duplicateRAM()
	fmt.Println(mypc)
}
