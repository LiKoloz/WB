package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

func main() {

	var wg sync.WaitGroup

	wg.Add(6)
	// Выход по условию
	go func(num int) {
		defer wg.Done()
		for i := range num {
			if i == 5 {
				return
			}
			fmt.Println(i * i)
		}
	}(7)

	ch := make(chan bool)

	// Отмена через канал
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("End chan goruting")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	// Выход по контексту
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("End context goruting")
				return
			default:

				fmt.Println("Горутина context работает")
				time.Sleep(1 * time.Second)
			}

		}
	}()
	time.Sleep(3 * time.Second)
	ch <- true
	cancel()

	// Выход с  помощью runtime.Goexit()
	go func(num int) {
		defer wg.Done()
		for i := range num {
			if i == 5 {
				runtime.Goexit()
			}
			fmt.Println(i * i)
		}
	}(7)

	//Выход с помощью os.Exit(0)
	go func(num int) {
		defer wg.Done()
		for i := range num {
			if i == 5 {
				return
			}
			fmt.Println(i * i)
		}
	}(7)

	// Выход по таймауту
	timeout := time.After(3 * time.Second)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-timeout:
				fmt.Println("Горутина остановлена по таймауту")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wg.Wait()
}
