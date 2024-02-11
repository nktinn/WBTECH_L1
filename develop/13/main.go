package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func AriphthemicSwap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func BitSwap(a, b int) (int, int) { // XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	a, b := 10, 12
	fmt.Printf("Before a: %d, b: %d\n", a, b)
	a, b = AriphthemicSwap(a, b)
	fmt.Printf("AriphthemicSwap -> a: %d, b: %d\n", a, b)
	a, b = BitSwap(a, b)
	fmt.Printf("BitwiseSwap -> a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("EasySwap -> a: %d, b: %d\n", a, b)
}
