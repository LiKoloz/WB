package main

import (
	"math/big"
	"testing"
)

func TestOperation(t *testing.T) {
	var arg big.Int = *big.NewInt(25e17)
	var exp *big.Float = new(big.Float).Mul(new(big.Float).SetInt(&arg), new(big.Float).SetInt(&arg))
	op := func(i1, i2 big.Int) big.Float {
		return *new(big.Float).Mul(new(big.Float).SetInt(&i1), new(big.Float).SetInt(&i2))
	}
	r := operation(arg, arg, op)

	if r.Cmp(exp) != 0 {
		t.Errorf("Bad result")
	}
}
