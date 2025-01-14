/**
https://stepik.org/lesson/228263/step/5?unit=200796

Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

Формат входных данных
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0
(само число 0 в последовательность не входит, а служит как признак ее окончания).

Формат выходных данных
Выведите ответ на задачу.

Sample Input:
1
3
3
1
0

Sample Output:
2
*/

package main

import "fmt"

func main() {
	var a int
	var max int = 0
	var cnt int = 0

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > max {
			max = a
			cnt = 0
		}
		if a == max {
			cnt++
		}
	}

	fmt.Println(cnt)
}
