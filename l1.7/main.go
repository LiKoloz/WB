package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		mut      sync.Mutex
		wg       sync.WaitGroup
		mapa     = make(map[int]int)
		counters sync.Map
	)
	wg.Add(5)
	for i := range 5 {
		go func() {
			defer wg.Done()
			mut.Lock()
			time.Sleep(1 * time.Second)
			mapa[i] = i * i
			mut.Unlock()
		}()
	}

	wg.Wait()
	for j, v := range mapa {
		fmt.Println(j, " - ", v)
	}

	for i := range 5 {
		go func() {
			time.Sleep(1 * time.Second)
			counters.Store(i, i*i)
		}()
	}
}
