/**
1.12 Массивы и срезы

https://stepik.org/lesson/228265/step/14?unit=200798
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:
2
3
56
45
21

Sample Output:
56
*/

package main

import "fmt"

func main() {
	array := [5]int{}

	var a, max int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// your code
	max = array[0]
	for i := 1; i < 5; i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	fmt.Println(max)
}
