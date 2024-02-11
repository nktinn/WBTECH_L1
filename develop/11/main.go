package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	set1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	set2 := map[int]int{3: 3, 4: 4, 5: 5, 6: 6, 7: 7}
	twoSets := make(map[int]int)

	for key, value := range set1 {
		if val, ok := set2[key]; ok && val == value { // Если ключ есть в обоих множествах и значения по ключу равны,
			twoSets[key] = value // то добавляем его в множество пересечений
		}
	}

	fmt.Printf("Пересечение множеств: %d", twoSets)
}
