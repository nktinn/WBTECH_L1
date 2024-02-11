package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func CustomMap(arr []string) map[string]bool {
	set := make(map[string]bool)
	for _, v := range arr {
		set[v] = true
	}
	return set
}

func CustomSet(arr []string) mapset.Set {
	customSet := mapset.NewSet()
	for _, v := range arr {
		customSet.Add(v)
	}
	return customSet
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	customMap := CustomMap(arr) // map
	customSet := CustomSet(arr) // set

	fmt.Printf("%v\n", customMap)
	fmt.Printf("%v\n", customSet)

	fmt.Printf("\"cat\" in map: %v\n", customMap["cat"])          // Получение значения по ключу
	fmt.Printf("\"cat\" in set: %v\n", customSet.Contains("cat")) // Получение информации о наличии элемента
}
