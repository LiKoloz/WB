package main

import (
	"fmt"
	"sync"
	"time"
)

// ConcMap - асинхронная MAP
type ConcMap struct {
	mutex sync.Mutex
	mapa  map[int]int
}

// Add - добавление в ассинхронную MAP
func (concmap *ConcMap) Add(i int) {
	concmap.mutex.Lock()
	if _, e := concmap.mapa[i]; e {
		panic("Int exist!")
	}
	concmap.mapa[i] = i * i
	time.Sleep(1 * time.Second)
	concmap.mutex.Unlock()
}
func main() {
	var (
		wg       sync.WaitGroup
		counters sync.Map
	)
	// Cинхронная MAP
	{
		wg.Add(5)
		concmap := ConcMap{
			mutex: sync.Mutex{},
			mapa:  make(map[int]int),
		}
		for i := range 5 {
			go func() {
				concmap.Add(i)
				defer wg.Done()
			}()
		}

		wg.Wait()
		for j, v := range concmap.mapa {
			fmt.Println(j, " - ", v)
		}
	}
	// Ассинхронная MAP
	{
		for i := range 5 {
			go func() {
				time.Sleep(1 * time.Second)
				counters.Store(i, i*i)
			}()
		}
	}
}
