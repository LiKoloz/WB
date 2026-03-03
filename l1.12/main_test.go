package main

import (
	"reflect"
	"testing"
)

func TestTakeSet(t *testing.T) {
	got := TakeSet([]string{"dog", "dog", "tree", "cat", "tree", "cat"})
	exp := []string{"dog", "tree", "cat"}

	if !reflect.DeepEqual(got, exp) {
		t.Error("Error")
	}
}
