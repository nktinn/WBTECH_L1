package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func CtxWorker(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("Waiting context to cancel...")
			time.Sleep(time.Second)
		case <-ctx.Done():
			fmt.Println("Worker stopped by cancelled context")
			return
		}
	}
}

func ChanWorker(done chan struct{}) {
	for {
		select {
		default:
			fmt.Println("Waiting channel to close...")
			time.Sleep(time.Second)
		case <-done:
			fmt.Println("Worker stopped by closed channel")
			return
		}
	}
}

func main() {
	// Остановка горутины с использованием контекста
	ctx, cancel := context.WithCancel(context.Background())
	go CtxWorker(ctx)
	time.Sleep(time.Second)
	cancel()

	// Остановка горутины с использованием канала
	done := make(chan struct{})
	go ChanWorker(done)
	time.Sleep(time.Second)
	close(done)

	time.Sleep(time.Second)
}
