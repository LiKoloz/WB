package main

import "strings"

func isStrUnique(str string) bool {
	m := make(map[string]struct{})

	for i := 0; i < len(str); i++ {
		j := strings.ToLower(string(str[i]))
		if _, e := m[j]; e {
			return false
		} else {
			m[j] = struct{}{}
		}
	}

	return true
}
