package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sleep(10, &wg)
	wg.Wait()
}

func sleep(n int, wg *sync.WaitGroup) {
	select {
	case <-time.After(time.Second * time.Duration(n)):
		fmt.Println("End sleep")
	}
	wg.Done()
}
