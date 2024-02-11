package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(600 * time.Millisecond)
			ch <- i
		}
	}()

	go func() {
		for {
			value := <-ch
			fmt.Println("Прочитано:", value)
		}
	}()

	time.Sleep(5 * time.Second) // Таймер, который блокирует завершение основного потока на необходимое время
	fmt.Println("Время вышло.")
	close(ch)
}
