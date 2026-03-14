package main

import (
	"reflect"
	"testing"
)

func TestGetAnnagramms(t *testing.T) {
	var (
		got = getAnagramms([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"})
		exp = map[string][]string{
			"листок": []string{"листок", "слиток", "столик"},
			"пятак":  []string{"пятак", "пятка", "тяпка"},
		}
	)

	if !reflect.DeepEqual(got, exp) {
		t.Error("Error")
	}
}
