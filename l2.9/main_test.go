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

func TestUnpackingString3(t *testing.T) {
	a := "qwe\\4\\5"
	e := "qwe45"

	r, err := unpacking_string(a)

	if err != nil && r != e {
		t.Error("error")
	}
}

func TestUnpackingString4(t *testing.T) {
	a := "qwe\\45"
	e := "qwe44444"

	r, err := unpacking_string(a)

	if err != nil && r != e {
		t.Error("error")
	}
}
