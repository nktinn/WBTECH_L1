package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Sleep1(ms int) {
	<-time.After(time.Millisecond * time.Duration(ms)) // After ждет указанное время и возвращает канал с текущим временем
}

func Sleep2(ms int) {
	endTime := time.Now().UnixNano() + int64(ms)*int64(time.Millisecond)
	for time.Now().UnixNano() < endTime {
		// do nothing
	}
}

func main() {
	fmt.Println("Program started")
	Sleep1(2000)
	fmt.Println("Sleep1 finished")
	Sleep2(2000)
	fmt.Println("Sleep2 finished")
}
