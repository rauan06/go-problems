# Объяснение / Решение

Слайсы в го выглядят вот так
```go
type slice struct {
    bucket *unsafe.Pointer
    len()
    cap()
}
```

И когда мы копируем значение а
```go
a := []int{1, 2, 3}
b := a
```

Мы на самом деле копируем пойнтер на бакет. Поэтому, когда мы изменяем индекс 0 у слайса b,
b[0] = 20, то же самое происходит и у слайса а. Но append() работает по другому, append() не
меняет значение по адресу или индексу, он возвращает новый массив.
```go
// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
//
// As a special case, it is legal to append a string to a byte slice, like this:
//
//	slice = append([]byte("hello "), "world"...)
func append(slice []Type, elems ...Type) []Type
```