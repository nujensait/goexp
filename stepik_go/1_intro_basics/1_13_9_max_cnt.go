/**
1.13 Решение задач
https://stepik.org/lesson/229320/step/9?unit=201906

Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные
Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4

Sample Output:
1
*/

package main

import (
	"fmt"
)

func main() {
	var n, i, val int
	var a []int // array

	fmt.Scan(&n)

	var min = -1
	var cnt = 0

	for i = 0; i < n; i++ {
		fmt.Scan(&val)
		a = append(a, val)
		if val < min || min == -1 {
			min = val
		}
	}
	//fmt.Print("min: ")
	//fmt.Println(min)

	// count mins
	for i = 0; i < n; i++ {
		if a[i] == min {
			cnt++
		}
	}

	//fmt.Print("res: ")
	fmt.Println(cnt)
}
