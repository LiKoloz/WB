package main

import "fmt"

func main() {
	// Установка первого бита в 1: 110 | 001 = 111
	a := 6 | 1
	fmt.Println(a)

	// Установка первого бита в 0: 111 | 110 = 110
	b := 7 & 6
	fmt.Println(b)
}
