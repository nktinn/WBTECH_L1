package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	i := 2
	slice = append(slice[:i], slice[i+1:]...) // :i не входит в итоговый слайс
	fmt.Println(slice)
}
