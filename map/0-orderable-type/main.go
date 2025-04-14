// //////////////////////////////////////
//
// Какие строки кода приведут к панике,
// а какие будут выполнены без ошибок?
//

package main

import "fmt"

func main() {
	WriteResponse("my numbers", 12345)
	WriteResponse("🪲👍", "emoji")
	WriteResponse([]int{1, 2, 3}, "массив чисел")
	WriteResponse(map[int]string{}, "словарь")
	WriteResponse(make(chan int), "канал")
	WriteResponse([]rune{'🙌', '👌'}, "руны")
}

func WriteResponse(key interface{}, value interface{}) {
	defer Recover()

	var myMap = map[interface{}]interface{}{}

	myMap[key] = value

	fmt.Println(myMap)
}
