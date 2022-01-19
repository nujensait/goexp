/**
https://stepik.org/lesson/229320/step/6?unit=201906

Входные данные
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.

Выходные данные
Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует". Строку выводите без кавычек.

Sample Input:
4 5 6

Sample Output:
Существует
*/

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var res bool

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	// if not enough length of two sides...
	if a+b < c || a+c < b || b+c < a {
		res = false
	} else {
		res = true
	}

	if res {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
