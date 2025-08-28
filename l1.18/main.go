package main

import "sync"

type Counter struct {
	counter int
	mutex   sync.Mutex
}

func (c *Counter) add() {
	c.mutex.Lock()
	c.counter++
	c.mutex.Unlock()
}
