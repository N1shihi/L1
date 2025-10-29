package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("2142141244214124214", 10)
	b.SetString("42242442424245363421424124214", 10)

	sum := new(big.Int)
	diff := new(big.Int)
	prod := new(big.Int)
	quot := new(big.Int)
	rem := new(big.Int)

	sum.Add(a, b)
	diff.Sub(a, b)
	prod.Mul(a, b)
	quot.Div(b, a)
	rem.Mod(b, a)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", prod)
	fmt.Println("b / a =", quot)
	fmt.Println("b % a =", rem)
}
