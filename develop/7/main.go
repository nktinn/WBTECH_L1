package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	var sm sync.Mutex
	var wg sync.WaitGroup

	m := make(map[int]int)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Lock()
			m[i] = i
			sm.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(m[1], m[5], m[10], m[19])
}
