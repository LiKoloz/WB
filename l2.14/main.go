package main

import (
	"fmt"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{}

func main() {
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		result := make(chan interface{}, 1)
		for _, c := range channels {
			go func() {

				if v, o2 := <-c; o2 {
					result <- v
				} else {
					result <- v
					close(result)
				}

			}()
		}
		return result
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	var a, b = <-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println(a, " ", b)
	fmt.Printf("done after %v", time.Since(start))
}
