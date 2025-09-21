package main

import "testing"

func TestIsUniqueStr1(t *testing.T) {
	a := "abcd"
	e := true

	r := isStrUnique(a)

	if e != r {
		t.Error("Bad result")
	}
}

func TestIsUniqueStr2(t *testing.T) {
	a := "abCdefAaf"
	e := false

	r := isStrUnique(a)

	if e != r {
		t.Error("Bad result")
	}
}

func TestIsUniqueStr3(t *testing.T) {
	a := "aabcd"
	e := false

	r := isStrUnique(a)

	if e != r {
		t.Error("Bad result")
	}
}
