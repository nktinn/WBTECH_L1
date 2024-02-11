package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа
должна выводить итоговое значение счетчика
*/

type Counter struct {
	value int
	sm    sync.Mutex
}

func (c *Counter) Inc() {
	c.sm.Lock()
	c.value++
	c.sm.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	count := 23 // Количество конкурентных методов инкремента счетчика
	counter := Counter{}

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Counter value: %d\n", counter.Value())
}
