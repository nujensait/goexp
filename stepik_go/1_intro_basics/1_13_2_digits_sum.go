/**
1.13 Решение задач

Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:
745

Sample Output:
16
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var i, sum int

	fmt.Scan(&a)

	for i = 0; i < len(a); i++ {
		val, err := strconv.Atoi(string(a[i]))
		if err == nil {
			sum += val
		}
	}

	fmt.Print(sum)
}
