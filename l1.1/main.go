package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHello() {
	fmt.Println("Hello ", h.name)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human{
			name: "Ilya",
			age:  22,
		},
	}
	a.SayHello()
}
