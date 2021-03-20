package main

import (
	pk "curso-golang-platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	var otherCar pk.CarPublic
	fmt.Println(otherCar)

	pk.PrintMessage("hola platzi")
}
