package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBitXOR(number int64, pos uint) int64 {
	number ^= 1 << pos // XOR между значением и битом, сдвинутым на необходимую позицию
	return number
}

func main() {
	var number int64 = 128
	fmt.Printf("До: %d\n", number)

	number = setBitXOR(number, 7)
	fmt.Printf("После: %d\n", number)
}
