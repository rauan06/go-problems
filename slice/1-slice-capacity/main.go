//////////////////////////
//
// Что выведет программа?
//

package main

import "fmt"

func fn(s []int) {
	s[2] = 5
	s = append(s, 6)
	s = append(s, 7)
	s[0] = 10
}

func main() {
	s := make([]int, 0, 5)
	for i := 0; i < 4; i++ {
		s = append(s, i)
	}

	fn(s[:3])

	fmt.Println(s)
}
