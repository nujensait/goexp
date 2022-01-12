/**
1.12 Массивы и срезы
https://stepik.org/lesson/228265/step/15?unit=200798

Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные
Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100).
Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.

Выходные данные
Необходимо вывести все элементы массива с чётными индексами.

Sample Input:
6
4 5 3 4 2 3

Sample Output:
4 3 2

(!) Usage of slices here is required
*/

package main

import "fmt"

func main() {
	var a, n, i, j int
	var res []int
	j = 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&a)
		if i%2 == 0 {
			res = append(res, a)
			j++
		}
	}

	for i = 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}
