// ///////////////////////
//
// Вспомоггательный файл
//

package main

import "fmt"

func Recover() {
	if v := recover(); v != nil {
		fmt.Println("\033[31m"+"Ошибка: ", v, "\033[0m")
	}
}
