package main

import "testing"

func TestRever(t *testing.T) {
	arg := "главрыба"
	exp := "абырвалг"

	v := reverseStr(arg)

	if v != exp {
		t.Errorf("Bad result")
	}
}
