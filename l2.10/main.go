package main

import (
	"fmt"
	"strconv"
	"strings"
)

// SortByColmn - Сортировка по колоке
func SortByColmn(strs []string, k int) {
	for i := 0; i < len(strs); i++ {
		for j := 1; j < len(strs); j++ {
			if rune(strs[i][k]) > rune(strs[j][k]) {
				tmp := strs[i]
				strs[i] = strs[j]
				strs[j] = tmp
			}
		}
	}
	fmt.Println(strs)
}

// GetHashCode - Получение хеша из строки
func GetHashCode(str string) rune {
	result := rune(0)
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		result += (rune(str[i]) % 2)
	}
	return result
}

// SortByHash - Сортировка по хешу в строке
func SortByHash(strs []string) {
	for i := 0; i < len(strs); i++ {
		for j := 1; j < len(strs); j++ {
			if GetHashCode(strs[i]) > GetHashCode(strs[j]) {
				tmp := strs[i]
				strs[i] = strs[j]
				strs[j] = tmp
			}
		}
	}
	fmt.Println(strs)
}

// Reverse - Переворачивание строки
func Reverse(strs []string) {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

}

// UniqueStrs - Удаление дубликатов
func UniqueStrs(strs []string) {
	ustrs := make(map[string]struct{})
	for i := 0; i < len(strs); i++ {
		if _, e := ustrs[strs[i]]; !e {
			ustrs[strs[i]] = struct{}{}
		}
	}
	strs = make([]string, 0)
	for i := range ustrs {
		strs = append(strs, i)
	}
}

func main() {
	var (
		s    string
		strs = make([]string, 0)
	)
	fmt.Scan(&s)
	strs = strings.Split(s, " ")
	last := strs[len(strs)-1]
	if strings.HasPrefix(last, "-") {
		for i := 1; i < len(last); i++ {
			switch last[i] {
			case 'n':
				SortByHash(strs)
			case 'r':
				Reverse(strs)
			case 'u':
				UniqueStrs(strs)
			}
		}
	} else {
		col, e := strconv.ParseInt(last, 10, 0)
		if e != nil {
			return
		}
		last := strs[len(strs)-2]
		for i := 1; i < len(last); i++ {
			switch last[i] {
			case 'n':
				SortByHash(strs)
			case 'r':
				Reverse(strs)
			case 'u':
				UniqueStrs(strs)
			case 'k':
				SortByColmn(strs, int(col-1))
			}
		}
	}
	fmt.Println(strs)
}
