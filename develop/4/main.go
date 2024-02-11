package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

/*1 Способ*/ /*
func main() {
	workerCount := 5
	ch := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case value := <-ch:
					println("Worker", i+1, ":", value)
				case <-done:
					println("Worker", i+1, "stopped")
					return
				}
			}
		}(i)
	}

	// Завершение по нажатию Ctrl+C
	quit := make(chan os.Signal, 1)                      // Создание канала типа сигнал с буффером 1
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // Подписка на сигналы завершения программы, syscall.SIGINT == "Ctrl+C"
	<-quit                                               // Ожидание сигнала и последующее завершение программы

	close(done) // Закрытие канала для завершения работы всех воркеров
	wg.Wait()   // Ожидание завершения работы всех воркеров
	close(ch)   // Закрытие канала чтения-записи
}*/

/*2 Способ*/
func main() {
	workerCount := 5
	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case value := <-ch:
					fmt.Println("Worker", i+1, ":", value)
				case <-ctx.Done():
					fmt.Println("Worker", i+1, "stopped")
					return
				}
			}
		}(i, ctx)
	}

	// Завершение по нажатию Ctrl+C
	quit := make(chan os.Signal, 1)                      // Создание канала с буффером 1
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // Подписка на сигналы завершения программы, syscall.SIGINT == "Ctrl+C"
	<-quit                                               // Ожидание сигнала и последующее завершение программы

	cancel() // Завершение работы всех воркеров

	wg.Wait()
	close(ch)
}
