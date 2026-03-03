package main

import (
	"fmt"
	"sync"
)

func main() {
	mas := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	wg.Add(5)
	for _, v := range mas {
		go func() {
			fmt.Println(v * v)
			wg.Done()
		}()
	}
	wg.Wait()
}
