package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func Reverse(str string) string {
	runed := []rune(str) // Использование рун для корректной работы с юникодом
	length := len(runed)
	reversed := make([]rune, length)
	for i, s := range runed {
		reversed[length-1-i] = s
	}
	return string(reversed)
}

func main() {
	str := "главрыба"
	reversed := Reverse(str)
	fmt.Printf("%s —> %s\n", str, reversed)

	str = "主要魚類😊"
	reversed = Reverse(str)
	fmt.Printf("%s —> %s\n", str, reversed)
}
