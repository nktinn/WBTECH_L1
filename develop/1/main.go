package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greet() string {
	return "Hello, " + h.Name + "!"
}

func (h *Human) SayBye() string {
	return "Bye, " + h.Name + "!"
}

type Action struct {
	Human
	Profession string
}

func (a *Action) Greet() string {
	return "Hello, " + a.Name + ". Are you a " + a.Profession + "?"
}

func main() {
	h := Human{"Bob", 12}
	fmt.Println(h.Greet())

	a := Action{Human{"John", 25}, "go developer"}

	// Использование метода структуры Human с явным указанием
	fmt.Println(a.Human.Greet())
	// Использование метода анонимного типа Human
	fmt.Println(a.SayBye())

	// Различные варианты использования полей встроенной структуры
	fmt.Printf("%s, %d years old\n", a.Name, a.Human.Age)

	// Переопределенный метод структуры Action
	fmt.Println(a.Greet())
}
