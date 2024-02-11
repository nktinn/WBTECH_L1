package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат
операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {

	array := []int{2, 4, 6, 8, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	// Запись данных в первый канал
	go func() {
		defer close(ch1)
		defer wg.Done()
		for _, v := range array {
			ch1 <- v
		}
	}()

	// Умножение чисел на 2
	go func() {
		defer close(ch2)
		defer wg.Done()
		for v := range ch1 {
			ch2 <- v * 2
		}
	}()

	// Вывод данных из второго канала
	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
