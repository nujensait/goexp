/**
1.13 Решение задач
https://stepik.org/lesson/229320/step/8?unit=201906

Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:
5
1
8
100
0
12

Sample Output:
1
*/

package main

import "fmt"

func main() {
	var a, n, i int
	var cnt = 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			cnt++
		}
	}

	fmt.Print(cnt)
}
