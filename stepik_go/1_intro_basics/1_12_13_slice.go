/**
1.12 Массивы и срезы
https://stepik.org/lesson/228265/step/13?unit=200798

Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4),
а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/

package main

import "fmt"

func main() {
	var a, n, i int
	var arr [100]int

	fmt.Scan(&n)

	for i = 0; i <= n; i++ {
		fmt.Scan(&a)
		arr[i] = a
	}

	fmt.Print(arr[3])
}
