// ////////////////////////
//
// Что выведет программа?
//

package main

import "fmt"

func main() {
	var myMap = make(map[chan int]string)
	key1 := make(chan int)

	myMap[key1] = "my value"
	fmt.Println(myMap[key1])

	go func() {
		key1 <- 1
		key1 <- 2
		key1 <- 3
	}()

	<-key1
	fmt.Println(myMap[key1])

	close(key1)
	fmt.Println(myMap[key1])
}
