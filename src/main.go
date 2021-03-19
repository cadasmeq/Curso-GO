package main

import "fmt"

func main() {
	// Array
	var array [5]int
	array[0] = 1
	array[1] = 2

	fmt.Println(len(array), cap(array))

	// Slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, len(slice), cap(slice))

	// Sliceing
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

}
