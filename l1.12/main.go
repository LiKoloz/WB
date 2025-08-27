package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	new_strs := []string{}
	m := make(map[string]struct{})

	for _, s := range strs {
		if _, e := m[s]; !e {
			new_strs = append(new_strs, s)
			m[s] = struct{}{}
		}
	}

	fmt.Println(new_strs)
}
