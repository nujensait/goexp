/*
https://stepik.org/lesson/228838/step/11

Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

package main

import (
	"fmt"
)

func main() {
	var cnt, sum = sumInt(1, 2, 3, 4, 5)
	fmt.Println(cnt, sum)
}

func sumInt(a ...int) (int, int) {
	sum := 0
	for _, elem := range a {
		sum += elem
	}
	return len(a), sum
}
