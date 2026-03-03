package main

// SetBit - устанавливает i-й бит числа n в 1
func SetBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// ClearBit -  устанавливает i-й бит числа n в 0
func ClearBit(n int64, i uint) int64 {
	return n & ^(1 << i)
}
