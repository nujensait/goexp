/**
@source https://stepik.org/lesson/228838/step/9?auth=login&unit=201372

Последовательность Фибоначчи определена следующим образом:
φ1=1, φ2=1, φn=φn-1 + φn-2 при n > 1.

Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...

Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные
Вводится одно число n.

Выходные данные
Необходимо вывести  значение φn.

Sample Input:
4

Sample Output:
3
*/

package main

import (
	"fmt"
)

func main() {
	var fib = fibonacci(4)
	fmt.Println(fib)
}

func fibonacci(n int) int {
	var a [1000]int
	a[1] = 1
	a[2] = 1
	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
		//fmt.Println(a[i])
		//fmt.Println(" ")
	}
	return a[n]
}
