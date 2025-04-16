// ////////////////////////
//
// Что выведет программа?
//

package main

import "fmt"

func main() {
	var s = "🌍Hello!"

	fmt.Printf("Длина строки: %d\n", len(s))

	fmt.Print("Цикл по байтам: ")
	for i := range s {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

	fmt.Print("Цикл по рунам: ")
	for _, r := range s {
		fmt.Print(string(r))
	}
	fmt.Println()
}
