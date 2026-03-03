package main

import (
	"reflect"
	"testing"
)

func TestCrossingSlices(t *testing.T) {
	got := CrossingSlices([]int{1, 2, 3}, []int{2, 3, 4})
	exp := []int{2, 3}

	if !reflect.DeepEqual(got, exp) {
		t.Error("Error")
	}
}
