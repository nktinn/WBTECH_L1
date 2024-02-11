package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/

func Reverse(str string) string {
	s := strings.Fields(str) // Разбиваем строку на слова по пробелам
	length := len(s)
	reverse := make([]string, length)

	for i, value := range s {
		reverse[length-1-i] = value
	}
	return strings.Join(reverse, " ") // Объединяем слова в строку через пробел
}

func main() {
	str := "snow dog sun"
	reversed := Reverse(str)
	fmt.Printf("%s —> %s\n", str, reversed)

	str = "снег собака солнце"
	reversed = Reverse(str)
	fmt.Printf("%s —> %s\n", str, reversed)
}
