package main

import (
	"fmt"
)

func reverseWordStr(str string) string {
	mas := []rune(str)
	index := len(mas)
	right := index - 1
	left := right - 1
	flag := false

	for {
		if left == 0 {
			mas = append(mas, mas[left:right]...)
			break
		}
		if mas[left] != ' ' {
			left--
		} else {
			if flag == true {
				mas = append(mas, mas[left:right]...)
				if mas[len(mas)-1] != ' ' {
					mas = append(mas, ' ')
				}
				right = left
				left--
			} else {
				flag = true
				right = left
				index = left
				left--
			}
		}
	}
	mas = mas[index+1:]
	return string(mas)
}

func main() {
	fmt.Println(reverseWordStr("snow dog sun"))
	fmt.Println(reverseWordStr("tree bee free"))
}
