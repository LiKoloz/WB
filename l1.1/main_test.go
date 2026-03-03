package main

import (
	"reflect"
	"testing"
)

func TestMain(T *testing.T) {
	var a = Action{
		Human{
			name: "Ilya",
			age:  22,
		},
	}
	got := "Hello Ilya"
	exp := a.GetHello()

	if !reflect.DeepEqual(got, exp) {
		T.Error("Error")
	}
}
