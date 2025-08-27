package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer close(ch)
		defer wg.Done()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		mas := [5]int{r1.Intn(100), r1.Intn(100), r1.Intn(100), r1.Intn(100), r1.Intn(100)}

		for _, v := range mas {
			ch <- v
		}
	}()
	go func() {
		defer close(ch2)
		defer wg.Done()
		for v := range ch {
			ch2 <- v * v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
