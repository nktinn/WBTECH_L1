package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func Unique(str string) bool {
	str = strings.ToLower(str)
	s := []rune(str)
	unique := make(map[rune]bool)
	for _, r := range s {
		if _, ok := unique[r]; ok {
			return false
		}
		unique[r] = true
	}
	return true
}

func main() {
	str := "abcd"
	fmt.Println("Every char from", str, "is unique -", Unique(str))

	str = "abCdefA"
	fmt.Println("Every char from", str, "is unique -", Unique(str))

	str = "aabcd"
	fmt.Println("Every char from", str, "is unique -", Unique(str))
}
