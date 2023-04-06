package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(12345678)
	b := big.NewFloat(56473892)

	sum := new(big.Float).Add(a, b) // Метод для сложения big.Floats
	sub := new(big.Float).Sub(a, b) // Метод для вычитания big.Floats
	rem := new(big.Float).Quo(a, b) // Метод для деления big.Floats
	mul := new(big.Float).Mul(a, b) // Метод для умножения big.Float

	fmt.Printf(
		"Sum: %1.f\nSub: %1.f\nRem: %1.3f\nMul: %1.f\n",
		sum, sub, rem, mul,
	)
}
