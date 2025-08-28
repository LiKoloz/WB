package main

import (
	"fmt"
	"math/rand"
	"time"
)

// quickSort реализует алгоритм быстрой сортировки для среза целых чисел
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pivot := arr[r1.Intn(len(arr)-1)]

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		quickSort(arr[:right+1])
	}
	if left < len(arr) {
		quickSort(arr[left:])
	}
}

func main() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6, 0}
	fmt.Println("До сортировки:", numbers)

	quickSort(numbers)
	fmt.Println("После сортировки:", numbers)

	numbers2 := []int{-5, 2, -8, 1, 9, -3, 7, 4, -6, 0}
	fmt.Println("\nДо сортировки:", numbers2)

	quickSort(numbers2)
	fmt.Println("После сортировки:", numbers2)
}
