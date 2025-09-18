package main

import (
	"fmt"
	"math/big"
)

func operation(a big.Int, b big.Int, c func(big.Int, big.Int) big.Float) big.Float {
	fmt.Println(c(a, b))
	return c(a, b)
}
