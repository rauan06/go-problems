// ////////////////////////
//
// Что выведет программа?
//

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a

	b[0] = 20
	b = append(b, 2)
	b = append(b, 3)

	fmt.Println(a)
	fmt.Println(b)
}
