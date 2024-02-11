Сколько существует способов задать переменную типа slice или map?

Ответ:
```go
slice := []int{1, 2, 3}
slice := make([]int, 3) // [0 0 0] с автоматической емкостью 3
slice := make([]int, 3, 5) // [0 0 0] с емокстью 5
slice2 := append(slice[:2], slice[3:]...)

mymap := map[string]int{"a": 1, "b": 2, "c": 3}
mymap := make(map[string]int) // пустая мапа
mymap := make(map[string]int, 5) // с емкостью 5
```