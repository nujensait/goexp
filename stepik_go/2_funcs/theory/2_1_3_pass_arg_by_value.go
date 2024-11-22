/*
В качестве аргументов при вызове функции можно передавать и значения переменных,
результаты операций или других функций,

(*) но при этом следует учитывать, что если аргументы в функцию
передаются по значению то они копируются:
@see https://stepik.org/lesson/228838/step/5?auth=login&unit=201372
*/

package main

import "fmt"

func main() {
	var a = 8
	fmt.Println("a before: ", a)

	increment(a)
	fmt.Println("a after: ", a)
}

func increment(x int) {
	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}
