/**
1.13 Решение задач
https://stepik.org/lesson/229320/step/16?unit=201906

Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/

package main

import (
	"fmt"
)

func main() {
	var a, n, res string
	var i int

	fmt.Scan(&a)
	fmt.Scan(&n)

	for i = 0; i < len(a); i++ {
		if string(a[i]) != n {
			res += string(a[i])
		}
	}

	fmt.Print(res)
}
