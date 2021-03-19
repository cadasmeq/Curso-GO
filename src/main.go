package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es Palindromo.")
	} else {
		fmt.Println("No es un palindromo.")
	}
}

func main() {
	m := make(map[string]int)
	m["newmeta"] = 15
	m["client"] = 30

	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}
	// Encontrar valor -> return zero value
	value, ok := m["Joseph"]
	fmt.Println(value, ok)

	value, ok = m["newmeta"]
	fmt.Println(value, ok)
}
