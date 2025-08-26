package main

import (
	"fmt"
)

func main() {
	var num int
	var ch = make(chan bool)
	fmt.Print("Введите количество горутин: ")
	fmt.Scan(&num)

	for i := range num {
		go func(i int, ch <-chan bool) {
			for {
				var a = <-ch

				fmt.Println("Горутина ", i, " получила данные из канала: ", a)
			}
		}(i, ch)
	}

	for {
		ch <- true
	}
}
