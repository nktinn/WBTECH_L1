package main

import "fmt"

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором
*/

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) DistanceTo(second *Point) int {
	dx := p.x - second.x
	dy := p.y - second.y
	return dx*dx + dy*dy
}

func DistanceBetween(first, second *Point) int {
	dx := first.x - second.x
	dy := first.y - second.y
	return dx*dx + dy*dy
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(2, 2)

	fmt.Println("Расстояние =", point1.DistanceTo(point2))       // Метод
	fmt.Println("Расстояние =", DistanceBetween(point1, point2)) // Функция
}
