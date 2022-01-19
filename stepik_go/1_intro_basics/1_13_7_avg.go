/**
Даны два числа. Найти их среднее арифметическое.
https://stepik.org/lesson/229320/step/7?unit=201906

Формат входных данных
На вход дается два целых положительных числа a и b.

Формат выходных данных
Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)

Sample Input 1:
3 5
Sample Output 1:
4

Sample Input 2:
277 154

Sample Output 2:
215.5
*/

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	c = (a + b) / 2

	fmt.Println(c)
}
