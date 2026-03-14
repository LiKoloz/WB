package main

import "testing"

func TestSetBit(t *testing.T) {
	got := SetBit(5, 1)
	var exp int64 = 7

	if got != exp {
		t.Errorf("Error")
	}
}

func TestClearBit(t *testing.T) {
	got := ClearBit(5, 0)
	var exp int64 = 4

	if got != exp {
		t.Errorf("Error")
	}
}
