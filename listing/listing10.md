Что выведет данная программа и почему?

```go
func update(p *int) {
    b := 2
    p = &b
}

func main() {
    var (
        a = 1
        p = &a
    )
    fmt.Println(*p)
    update(p)
    fmt.Println(*p)
}

```

Ответ:

1 1

Так происходит из-за того, что в функции update создается копия указателя p, который ссылается на другой адрес памяти. 
Поэтому изменения в функции update не влияют на переменную p в функции main.