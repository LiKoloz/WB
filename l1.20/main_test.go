package main

import "testing"

func TestReverseWordStr(t *testing.T) {
	arg := "snow dog sun"
	exp := "sun dog snow"

	r := reverseWordStr(arg)

	if r != exp {
		t.Errorf("Bad result")
	}
}

func TestReverseWordStr2(t *testing.T) {
	arg := "tree bee free"
	exp := "free bee tree"

	r := reverseWordStr(arg)

	if r != exp {
		t.Errorf("Bad result")
	}
}
