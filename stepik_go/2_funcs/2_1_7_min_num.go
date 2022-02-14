/**
@source https://stepik.org/lesson/228838/step/7?unit=201372
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
Входные данные
Вводится четыре числа.
Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4

Run example:
go run .\2_1_7_min_num.go
4 2 1 3
1
*/

package main

import "fmt"

func main() {
	var min = minimumFromFour()
	fmt.Println(min)
}

func Min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumFromFour() int {
	var a [4]int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i])
	}
	min := 0
	min = Min2(a[0], a[1])
	min = Min2(min, a[2])
	min = Min2(min, a[3])
	return min
}
