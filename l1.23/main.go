package main

func deleteElement(m []int, n int) []int {
	if n == len(m)-1 {
		m = m[:n]
	} else {
		m = append(m[:n], m[n+1:]...)
	}
	return m
}
