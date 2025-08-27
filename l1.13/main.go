package main

import "fmt"

func main() {
	// Использование функционала Go
	var a, b = 10, 20
	b, a = a, b

	fmt.Println(a, " ", b)

	//Использование сложения и вычитания
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, " ", b)

	// Использование XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, " ", b)
}
