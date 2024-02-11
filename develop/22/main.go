package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
*/

func Sum(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(a, b)
	return res
}

func Sub(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Div(a, b)
	return res
}

func Mul(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(a, b)
	return res
}

func Div(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Div(a, b)
	return res
}

func main() {
	var a, b big.Int
	a.Lsh(big.NewInt(2), 23)
	b.Lsh(big.NewInt(2), 21)

	fmt.Println("a =", a.String(), "b =", b.String())
	fmt.Println("a + b =", Sum(&a, &b).String())
	fmt.Println("a - b =", Sub(&a, &b).String())
	fmt.Println("a * b =", Mul(&a, &b).String())
	fmt.Println("a / b =", Div(&a, &b).String())
}
