package main

import (
	"fmt"
	"sort"
	"strings"
)

// getAnagramms - получение аннаграмм
func getAnagramms(strs []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		word = strings.ToLower(word)
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		key := string(runes)

		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string)
	for _, group := range groups {
		if len(group) < 2 {
			continue
		}
		sort.Strings(group)
		result[group[0]] = group
	}
	return result
}

func main() {
	fmt.Println(getAnagramms([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}))
}
