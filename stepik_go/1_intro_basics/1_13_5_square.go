/**
https://stepik.org/lesson/229320/step/5?unit=201906

Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника.

Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный".
Иначе вывести "Непрямоугольный"

Sample Input:
6 8 10

Sample Output:
Прямоугольный
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var res bool

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2) ||
		math.Pow(b, 2) == math.Pow(a, 2)+math.Pow(c, 2) ||
		math.Pow(c, 2) == math.Pow(a, 2)+math.Pow(b, 2) {
		res = true
	} else {
		res = false
	}

	if res {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
