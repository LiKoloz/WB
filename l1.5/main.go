package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения.
По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func main() {

	//Использование контекста, в горутинах подписываемся на контекст, вся логика в основном потоке
	// {
	// 	var (
	// 		ch = make(chan bool)
	// 		t  int
	// 		wg sync.WaitGroup
	// 	)

	// 	fmt.Println("Enter time: ")
	// 	fmt.Scan(&t)
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	// 	defer cancel()

	// 	wg.Add(2)
	// 	go func() {
	// 		defer wg.Done()
	// 		defer close(ch)
	// 		for {
	// 			select {
	// 			case <-ctx.Done():
	// 				fmt.Println("End of producer")
	// 				return
	// 			case ch <- true:
	// 				fmt.Println("Put to chan")
	// 			}
	// 		}
	// 	}()
	// 	go func() {
	// 		defer wg.Done()
	// 		for {
	// 			select {
	// 			case <-ctx.Done():
	// 				fmt.Println("End of consumer")
	// 				return
	// 			case v := <-ch:
	// 				fmt.Println("Get ", v)
	// 			}
	// 		}
	// 	}()

	// 	wg.Wait()
	// }

	//Использование таймера
	{
		var t int
		fmt.Println("Enter time: ")
		fmt.Scan(&t)

		ch := make(chan bool)
		var wg sync.WaitGroup

		timeout := time.After(time.Duration(t) * time.Second)

		wg.Add(2)

		go func() {
			defer wg.Done()
			defer close(ch)

			for {
				select {
				case <-timeout:
					fmt.Println("End of producer")
					return
				default:
					select {
					case ch <- true:
						fmt.Println("Put to chan")
					default:
					}
				}
			}
		}()

		go func() {
			defer wg.Done()

			for {
				select {
				case <-timeout:
					fmt.Println("End of consumer")
					return
				case v, ok := <-ch:
					if !ok {
						fmt.Println("Channel closed")
						return
					}
					fmt.Println("Get ", v)
				}
			}
		}()

		wg.Wait()
	}
}
