package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func someFunc() {
	v := createHugeString(1 << 10)
	justString := make([]rune, 100)
	copy(justString, v[:100]) // Копирование вместо присваивания, чтобы не ссылаться на большой массив и GC мог его удалить
}

func main() {
	someFunc()
}
