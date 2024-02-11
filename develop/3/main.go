package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

// GO 1.21. Для GO 1.22 замыкание переменных для горутин в цикле можно не использовать
// Приоритетным способом яляется использование атомиков, так как они работают быстрее и не требуют блокировки

/*1 Способ - Mutex*/ /*
func main() {
	array := []int{2, 4, 6, 8, 10}
	var sum int

	var wg sync.WaitGroup
	wg.Add(len(array))
	var sm sync.Mutex

	for _, value := range array {
		go func(value int) {
			defer wg.Done()
			sm.Lock() // Предотвращение состояния гонки
			sum += value * value
			sm.Unlock()
		}(value)
	}

	wg.Wait()

	fmt.Printf("sum = %d", sum)
}*/

/*2 Способ - запись в канал*/
/*
func main() {
	array := []int{2, 4, 6, 8, 10}
	var sum int
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(array))

	for _, value := range array {
		go func(value int) {
			defer wg.Done()
			ch <- value * value
		}(value)
	}

	// Использование горутины для предотвращения deadlock
	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		sum += value
	}

	fmt.Printf("sum = %d", sum)
}
*/

/*3 Способ - атомик*/
func main() {
	array := []int32{2, 4, 6, 8, 10}
	var sum int32

	var wg sync.WaitGroup
	wg.Add(len(array))

	for _, value := range array {
		go func(value int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, value*value)
		}(value)
	}

	wg.Wait()

	fmt.Printf("sum = %d", sum)
}
