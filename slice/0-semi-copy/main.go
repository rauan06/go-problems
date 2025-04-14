// ////////////////////////
//
// Что выведет программа?
//

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a

	b[0] = 10
	c := append(a, 4)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
