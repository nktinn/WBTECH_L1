package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right { // Пока левая граница меньше или равна правой границе (не перешли друг друга)
		mid := left + (right-left)/2

		if arr[mid] == target { // Элемент найден
			return mid
		}

		if arr[mid] > target { // Если элемент в середине больше целевого значения, то ищем в левой части
			right = mid - 1
		} else { // Если элемент в середине меньше целевого значения, то ищем в правой части
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	target := 2
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("%d найден. Индекс %d\n", target, index)
	} else {
		fmt.Printf("%d не найден в массиве\n", target)
	}
}
