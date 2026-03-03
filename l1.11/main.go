package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	c := []int{}
	m := make(map[int]struct{})

	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, e := m[v]; e {
			c = append(c, v)
		}
	}

	fmt.Println(c)
}

// CrossingSlices - получение пересечений
func CrossingSlices(a []int, b []int) []int {
	var (
		c = make([]int, 0)
		d = make(map[int]struct{})
	)
	for _, v := range a {
		d[v] = struct{}{}
	}
	for _, v := range b {
		if _, e := d[v]; e {
			c = append(c, v)
		}
	}
	return c
}
