package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Через канал, канал будет пустым уже после 1ой горутины, плохо масштабируется
	// {
	// 	sigChan := make(chan os.Signal, 1)

	// 	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// 	go func() {
	// 		<-sigChan
	// 		fmt.Println("Завершение работы...")
	// 		os.Exit(0)
	// 	}()

	// 	fmt.Println("Программа запущена. Нажмите Ctrl+C для завершения.")
	// 	for {
	// 		fmt.Println("Работаю...")
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }

	// Использование контекста, насколько мне известно, является стандартом для подобного рода операций, хорошо масштабируется
	{
		sigChan := make(chan os.Signal, 1)

		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		wg.Add(5)

		for i := range 5 {
			go func(i int) {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						fmt.Println("Завершение работы горутины: ", i)
						return
					default:
						fmt.Println("Горутина ", i, " работает!")
						time.Sleep(2 * time.Second)
					}
				}
			}(i)
		}
		<-sigChan
		cancel()
		fmt.Println("Завершение работы программы")

		wg.Wait()
	}
}
