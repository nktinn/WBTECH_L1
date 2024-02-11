package main

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет
их квадраты в stdout.
*/

import (
	"fmt"
	"sync"
)

// GO 1.21
// Захват переменной (замыкание)

func main() {
	array := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(array))

	for _, value := range array {
		go func(value int) {
			defer wg.Done()
			fmt.Printf("%d * %d = %d\n", value, value, value*value)
		}(value) // Захват переменной и использование её копии внутри горутины
	}

	wg.Wait()
}

// GO 1.22 (6.02.2024)
// Пофиксили

/*
func main() {
	array := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(array))

	for _, value := range array {
		go func(){
			defer wg.Done()
			fmt.Printf("%d * %d = %d\n", value, value, value*value)
		}()
	}

	wg.Wait()
}
*/
