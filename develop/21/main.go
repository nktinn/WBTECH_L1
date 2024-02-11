package main

import (
	"fmt"
	"time"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Transport interface {
	Drive()
}

type Buggy struct {
	brand string
}

func (b *Buggy) Drive() {
	println("Driving buggy now")
}

type Car struct {
	brand string
}

func (c *Car) Drive() {
	println("Driving car now")
}

type Traveller struct {
	name string
}

func (t *Traveller) Travel(transp Transport) {
	transp.Drive()
}

func main() {
	buggy := &Buggy{brand: "Meyers Manx"}
	car := &Car{brand: "Honda"}
	traveller := &Traveller{name: "Ivan"}

	traveller.Travel(car)

	time.Sleep(2 * time.Second)
	fmt.Println("Road is blocked")
	fmt.Println("Changing transport...")

	traveller.Travel(buggy)
}
