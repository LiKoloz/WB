package main

import (
	"testing"
)

func TestDeleteElement(t *testing.T) {
	arg := []int{1, 2, 3, 4, 5}
	exp := []int{1, 2, 4, 5}

	r := deleteElement(arg, 2)

	for i := 0; i < len(exp); i++ {
		if exp[i] != r[i] {
			t.Errorf("Bad result")
		}
	}
}
func TestDeleteElement2(t *testing.T) {
	arg := []int{1, 2, 3, 4, 5}
	exp := []int{1, 2, 3, 4}

	r := deleteElement(arg, 4)

	for i := 0; i < len(exp); i++ {
		if exp[i] != r[i] {
			t.Errorf("Bad result")
		}
	}
}
