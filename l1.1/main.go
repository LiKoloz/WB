package main

import "fmt"

// Human класс
type Human struct {
	name string
	age  int
}

//SayHello - вывод на консоль
func (h Human) SayHello() {
	fmt.Println("Hello ", h.name)
}

// GetHello - получение для теста
func (h Human) GetHello() string {
	return "Hello " + h.name
}

// Action - наследник
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
