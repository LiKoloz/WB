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

// TakeSet - создание уникального множетсва
func TakeSet(mas []string) []string {
	var (
		set = make([]string, 0)
		m   = make(map[string]struct{})
	)
	for _, s := range mas {
		if _, e := m[s]; !e {
			set = append(set, s)
			m[s] = struct{}{}
		}
	}
	return set
}
