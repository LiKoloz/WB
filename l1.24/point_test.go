package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
	a := NewPoint(3, 3)
	b := NewPoint(5, 5)

	exp := 2.8284271247461903

	r := a.Distance(b)

	if r != exp {
		t.Error("Bad result")
	}
}
