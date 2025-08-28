package main

import "fmt"

func binSearch(mas []int, num int, start int) int {
	if len(mas) == 0 {
		if mas[0] == num {
			return start
		} else {
			return -1
		}
	}

	if mas[len(mas)/2] == num {
		return start + len(mas)/2
	} else if num > mas[len(mas)/2] {
		return binSearch(mas[len(mas)/2+1:], num, start+len(mas)/2+1)
	} else {
		return binSearch(mas[:len(mas)/2], num, start)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(binSearch(a, 1, 0))
	fmt.Println(binSearch(a, 2, 0))
	fmt.Println(binSearch(a, 4, 0))
	fmt.Println(binSearch(a, 5, 0))
	fmt.Println(binSearch(a, 6, 0))
	fmt.Println(binSearch(a, 7, 0))
}
