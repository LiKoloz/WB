package main

import "testing"

func TestSetBit(t *testing.T) {
	got := SetBit(5, 1) // устанавливаем 1-й бит в 1: 0101 -> 0111 = 7
	var exp int64 = 7

	if got != exp {
		t.Errorf("SetBit(5, 1) = %d; want %d", got, exp)
	}
}

func TestClearBit(t *testing.T) {
	got := ClearBit(5, 0) // сбрасываем 1-й бит в 0: 0101 -> 0100 = 4
	var exp int64 = 4

	if got != exp {
		t.Errorf("Error")
	}
}
