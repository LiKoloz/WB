package main

import "testing"

func TestUnpackingString1(t *testing.T) {
	a := "adcd"
	e := "abcd"

	r, err := unpacking_string(a)

	if err != nil && r != e {
		t.Error("error")
	}
}

func TestUnpackingString2(t *testing.T) {
	a := "a4bc2d5e"
	e := "aaaabccddddde"

	r, err := unpacking_string(a)

	if err != nil && r != e {
		t.Error("error")
	}
}
